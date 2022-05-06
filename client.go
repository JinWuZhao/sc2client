package sc2client

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/JinWuZhao/sc2client/sc2proto"
)

type StepState struct {
	PlayerId      uint32
	Steps         uint32
	Rpc           *RpcClient
	ReceivedChats <-chan *sc2proto.ChatReceived
	Stop          chan<- error
}

type PlayerSetup struct {
	Type       sc2proto.PlayerType
	Race       sc2proto.Race
	Name       string
	Difficulty sc2proto.Difficulty
	AIBuild    sc2proto.AIBuild
	Step       func(ctx context.Context, state *StepState)
}

type Client struct {
	displayMode  int
	windowWidth  int
	windowHeight int
	windowX      int
	windowY      int
	rpcTimeout   time.Duration
	launcher     *Launcher
	conn         *Connection
	rpc          *RpcClient
	deferList    []func()
	stop         chan error
	playerId     uint32
	step         func(context.Context, *StepState)
}

func ClientDisplayModeOpts(displayMode int) func(*Client) {
	return func(client *Client) {
		client.displayMode = displayMode
	}
}

func ClientWindowOpts(windowWidth int, windowHeight int, windowX int, windowY int) func(*Client) {
	return func(client *Client) {
		client.windowWidth = windowWidth
		client.windowHeight = windowHeight
		client.windowX = windowX
		client.windowY = windowY
	}
}

func ClientRpcOpts(timeout time.Duration) func(*Client) {
	return func(client *Client) {
		client.rpcTimeout = timeout
	}
}

func NewClient(opts ...func(*Client)) *Client {
	client := &Client{
		displayMode:  0,
		windowWidth:  1024,
		windowHeight: 768,
		windowX:      100,
		windowY:      100,
		rpcTimeout:   30 * time.Second,
		stop:         make(chan error, 1),
	}
	for _, option := range opts {
		option(client)
	}
	return client
}

func (c *Client) Init(ctx context.Context) error {
	host, port, err := GetLocalAddress()
	if err != nil {
		return fmt.Errorf("GetLocalAddress() error: %w", err)
	}
	defer func() {
		if err != nil {
			c.Finalize()
		}
	}()

	c.launcher, err = NewLauncher(host, port, c.displayMode, c.windowWidth, c.windowHeight, c.windowX, c.windowY)
	if err != nil {
		return fmt.Errorf("NewLauncher() error: %w", err)
	}
	err = c.launcher.StartProcess(ctx)
	if err != nil {
		return fmt.Errorf("StartProcess() error: %w", err)
	}
	c.deferList = append(c.deferList, func() {
		_ = c.launcher.StopProcess()
	})
	log.Println("game started:")
	log.Println("process pid:", c.launcher.ProcessPid())
	log.Println("server listen on:", host+":"+strconv.Itoa(port))
	time.Sleep(10 * time.Second)

	c.conn, err = DialSC2(ctx, host, port)
	if err != nil {
		return fmt.Errorf("DialSC2() error: %w", err)
	}
	c.deferList = append(c.deferList, func() {
		c.conn.Close()
	})

	c.rpc = NewRpcClient(c.conn, c.rpcTimeout)
	pingRsp, err := c.rpc.Ping(ctx)
	if err != nil {
		return fmt.Errorf("c.rpc.Ping() error: %w", err)
	}
	log.Println("game version:", pingRsp.GetGameVersion())
	log.Println("base build:", pingRsp.GetBaseBuild())
	log.Println("data version", pingRsp.GetDataVersion())
	log.Println("data build:", pingRsp.GetDataBuild())
	c.deferList = append(c.deferList, func() {
		_, _ = c.rpc.Quit(context.Background(), &sc2proto.RequestQuit{})
		time.Sleep(10 * time.Second)
	})

	return nil
}

func (c *Client) Finalize() {
	for i := len(c.deferList) - 1; i >= 0; i-- {
		c.deferList[i]()
	}
	c.deferList = nil
}

func (c *Client) HostGame(ctx context.Context, portConfig *PortConfig, gameMap string, players []*PlayerSetup, disableFog bool) error {
	var playerSetups []*sc2proto.PlayerSetup
	for _, player := range players {
		playerSetups = append(playerSetups, &sc2proto.PlayerSetup{
			Type:       player.Type.Enum(),
			Race:       player.Race.Enum(),
			PlayerName: proto.String(player.Name),
			Difficulty: player.Difficulty.Enum(),
			AiBuild:    player.AIBuild.Enum(),
		})
	}
	createGameRsp, err := c.rpc.CreateGame(ctx, &sc2proto.RequestCreateGame{
		Map: &sc2proto.RequestCreateGame_LocalMap{
			LocalMap: &sc2proto.LocalMap{
				MapPath: proto.String(gameMap),
			},
		},
		PlayerSetup: playerSetups,
		DisableFog:  proto.Bool(disableFog),
		RandomSeed:  proto.Uint32(uint32(time.Now().Unix())),
		Realtime:    proto.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("c.rpc.CreateGame() error: %w", err)
	}
	if createGameRsp.GetError() > sc2proto.ResponseCreateGame_MissingMap {
		return fmt.Errorf("create game error: %s, %s",
			createGameRsp.GetError().String(), createGameRsp.GetErrorDetails())
	}

	joinGameRsp, err := c.rpc.JoinGame(ctx, &sc2proto.RequestJoinGame{
		Participation: &sc2proto.RequestJoinGame_Race{
			Race: players[0].Race,
		},
		ServerPorts: &sc2proto.PortSet{
			GamePort: proto.Int32(int32(portConfig.Servers[0])),
			BasePort: proto.Int32(int32(portConfig.Servers[1])),
		},
		ClientPorts: []*sc2proto.PortSet{
			{
				GamePort: proto.Int32(int32(portConfig.Players[0])),
				BasePort: proto.Int32(int32(portConfig.Players[1])),
			},
		},
		PlayerName: proto.String(players[0].Name),
	})
	if err != nil {
		return fmt.Errorf("c.rpc.JoinGame() error: %w", err)
	}
	if joinGameRsp.GetError() > sc2proto.ResponseJoinGame_MissingParticipation {
		return fmt.Errorf("participant join game error: %s, %s",
			joinGameRsp.GetError().String(), joinGameRsp.GetErrorDetails())
	}

	c.playerId = joinGameRsp.GetPlayerId()
	c.step = players[0].Step

	return nil
}

func (c *Client) JoinGame(ctx context.Context, portConfig *PortConfig, players []*PlayerSetup) error {
	joinGameRsp, err := c.rpc.JoinGame(ctx, &sc2proto.RequestJoinGame{
		Participation: &sc2proto.RequestJoinGame_Race{
			Race: players[1].Race,
		},
		ServerPorts: &sc2proto.PortSet{
			GamePort: proto.Int32(int32(portConfig.Servers[0])),
			BasePort: proto.Int32(int32(portConfig.Servers[1])),
		},
		ClientPorts: []*sc2proto.PortSet{
			{
				GamePort: proto.Int32(int32(portConfig.Players[0])),
				BasePort: proto.Int32(int32(portConfig.Players[1])),
			},
		},
		PlayerName: proto.String(players[1].Name),
	})
	if err != nil {
		return fmt.Errorf("c.rpc.JoinGame() error: %w", err)
	}
	if joinGameRsp.GetError() > sc2proto.ResponseJoinGame_MissingParticipation {
		return fmt.Errorf("participant join game error: %s, %s",
			joinGameRsp.GetError().String(), joinGameRsp.GetErrorDetails())
	}

	c.playerId = joinGameRsp.GetPlayerId()
	c.step = players[1].Step

	return nil
}

func (c *Client) StartGameLoop(ctx context.Context) {
	go func() {
		stopStep := make(chan error, 1)
		var prevStep uint32
		receivedChats := make(chan *sc2proto.ChatReceived, 100)
	gameLoop:
		for ctx.Err() == nil {
			resp, err := c.rpc.Observation(ctx, &sc2proto.RequestObservation{})
			if err != nil {
				c.stop <- fmt.Errorf("c.rpc.Observation() error: %w", err)
				break gameLoop
			}
			playerResult := resp.GetPlayerResult()
			if len(playerResult) > 0 {
				c.stop <- nil
				break gameLoop
			}
			if c.step != nil {
				for _, chat := range resp.GetChat() {
					select {
					case receivedChats <- chat:
					default:
						log.Println("[WARN] receivedChats overflowed:", chat.String())
					}
				}

				if resp.GetObservation().GetGameLoop() > prevStep {
					c.step(ctx, &StepState{
						PlayerId:      c.playerId,
						Steps:         resp.GetObservation().GetGameLoop(),
						Rpc:           c.rpc,
						ReceivedChats: receivedChats,
						Stop:          stopStep,
					})
					prevStep = resp.GetObservation().GetGameLoop()
					select {
					case err := <-stopStep:
						if err != nil {
							c.stop <- fmt.Errorf("game loop step stopped with error: %w", err)
						} else {
							c.stop <- nil
						}
						break gameLoop
					default:
					}
				}
			}
		}
	}()
}

func (c *Client) WaitGameEnd(ctx context.Context) error {
	defer func() {
		_, _ = c.rpc.LeaveGame(context.Background(), &sc2proto.RequestLeaveGame{})
	}()
	var retErr error
	select {
	case <-ctx.Done():
	case err := <-c.stop:
		if err != nil {
			retErr = fmt.Errorf("game loop end with error: %w", err)
		}
	}
	return retErr
}
