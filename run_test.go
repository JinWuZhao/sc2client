package sc2client

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/JinWuZhao/sc2client/sc2proto"
)

type DirectorAgent struct {
	rpc      *RpcClient
	playerId uint32
}

func (m *DirectorAgent) OnStart(playerId uint32, rpc *RpcClient) {
	log.Println("OnStart()")
	m.rpc = rpc
	m.playerId = playerId
}

func (m *DirectorAgent) OnStep(ctx context.Context, state *StepState) {
	if state.Steps%100 == 0 {
		_, err := m.rpc.Action(ctx, &sc2proto.RequestAction{
			Actions: []*sc2proto.Action{
				{
					ActionChat: &sc2proto.ActionChat{
						Channel: sc2proto.ActionChat_Team.Enum(),
						Message: proto.String(fmt.Sprintf("cmd-create-siege-tank 3 红-%d", state.Steps)),
					},
				},
				{
					ActionChat: &sc2proto.ActionChat{
						Channel: sc2proto.ActionChat_Team.Enum(),
						Message: proto.String(fmt.Sprintf("cmd-move-toward 红-%d %d %d", state.Steps, rand.Intn(90)-45, rand.Intn(100)*2)),
					},
				},
				{
					ActionChat: &sc2proto.ActionChat{
						Channel: sc2proto.ActionChat_Team.Enum(),
						Message: proto.String(fmt.Sprintf("cmd-create-siege-tank 4 蓝-%d", state.Steps)),
					},
				},
				{
					ActionChat: &sc2proto.ActionChat{
						Channel: sc2proto.ActionChat_Team.Enum(),
						Message: proto.String(fmt.Sprintf("cmd-move-toward 蓝-%d %d %d", state.Steps, rand.Intn(90)+135, rand.Intn(100)*2)),
					},
				},
			},
		})
		if err != nil {
			state.Stop <- fmt.Errorf("rpc.Action() error: %w", err)
			return
		}
	foreachChats:
		for {
			select {
			case msg := <-state.ReceivedChats:
				log.Printf("recive message from player %d: %s\n", msg.GetPlayerId(), msg.GetMessage())
			default:
				break foreachChats
			}
		}
	}
}

func (m *DirectorAgent) OnEnd(result sc2proto.Result) {
	log.Println("OnEnd():", result.String())
}

func TestRun_RunGame(t *testing.T) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	err := RunGame(ctx,
		[]GameMap{
			{Name: "StarArena.SC2Map"},
		},
		[]*PlayerSetup{
			{
				Type:       sc2proto.PlayerType_Participant,
				Race:       sc2proto.Race_Random,
				Name:       "Director",
				Difficulty: sc2proto.Difficulty_Easy,
				AIBuild:    sc2proto.AIBuild_RandomBuild,
				Agent:      new(DirectorAgent),
			},
			{
				Type:       sc2proto.PlayerType_Participant,
				Race:       sc2proto.Race_Random,
				Name:       "Audience",
				Difficulty: sc2proto.Difficulty_Easy,
				AIBuild:    sc2proto.AIBuild_RandomBuild,
			},
		},
		true)
	if err != nil {
		t.Error(err)
	}
}
