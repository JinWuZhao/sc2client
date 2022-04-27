package sc2client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jinwuzhao/sc2client/sc2proto"
)

type IDGenerator struct {
	id uint32
}

func (m *IDGenerator) Next() uint32 {
	return atomic.AddUint32(&m.id, 1)
}

type RpcClient struct {
	conn     *Connection
	idgen    IDGenerator
	respPool map[uint32]chan *sc2proto.Response
	mutex    sync.RWMutex
	timeout  time.Duration
}

func NewRpcClient(conn *Connection, timeout time.Duration) *RpcClient {
	rpc := &RpcClient{
		conn:     conn,
		respPool: make(map[uint32]chan *sc2proto.Response),
		timeout:  timeout,
	}
	go rpc.runMsgLoop()
	return rpc
}

func (c *RpcClient) runMsgLoop() {
	var err error
	for {
		resp := &sc2proto.Response{}
		err = c.conn.Read(context.Background(), resp)
		if err != nil {
			err = fmt.Errorf("c.conn.Read() error: %w", err)
			break
		}
		c.mutex.RLock()
		respChan, ok := c.respPool[resp.GetId()]
		if ok {
			select {
			case respChan <- resp:
			default:
			}
		}
		c.mutex.RUnlock()
	}
	log.Println("rpc client message loop stopped. reason:", err)
}

func (c *RpcClient) SendRequest(ctx context.Context, req *sc2proto.Request) (uint32, error) {
	reqID := c.idgen.Next()
	respChan := make(chan *sc2proto.Response, 1)
	c.mutex.Lock()
	c.respPool[reqID] = respChan
	c.mutex.Unlock()

	req.Id = &reqID
	err := c.conn.Write(ctx, req)
	if err != nil {
		c.mutex.Lock()
		delete(c.respPool, reqID)
		close(respChan)
		c.mutex.Unlock()
		return 0, fmt.Errorf("c.conn.Write() error: %w", err)
	}
	return reqID, nil
}

func (c *RpcClient) WaitForResponse(reqID uint32) (*sc2proto.Response, error) {
	c.mutex.RLock()
	respChan, ok := c.respPool[reqID]
	c.mutex.RUnlock()
	if !ok {
		return nil, fmt.Errorf("invalid request ID: %d", reqID)
	}
	defer func() {
		c.mutex.Lock()
		delete(c.respPool, reqID)
		close(respChan)
		c.mutex.Unlock()
	}()
	var resp *sc2proto.Response
	select {
	case resp = <-respChan:
		return resp, nil
	case <-time.After(c.timeout):
		return nil, fmt.Errorf("waiting for response timeout: %d", reqID)
	}
}

func (c *RpcClient) Ping(ctx context.Context) (*sc2proto.ResponsePing, error) {
	id, err := c.SendRequest(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_Ping{
			Ping: &sc2proto.RequestPing{},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("c.SendRequest() error: %w", err)
	}
	resp, err := c.WaitForResponse(id)
	if err != nil {
		return nil, fmt.Errorf("c.WaitForResponse() error: %w", err)
	}
	if len(resp.GetError()) > 0 {
		return nil, fmt.Errorf("sc2 client response error: %+v", resp.GetError())
	}
	return resp.GetPing(), nil
}

func (c *RpcClient) CreateGame(ctx context.Context, req *sc2proto.RequestCreateGame) (*sc2proto.ResponseCreateGame, error) {
	id, err := c.SendRequest(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_CreateGame{
			CreateGame: req,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("c.SendRequest() error: %w", err)
	}
	resp, err := c.WaitForResponse(id)
	if err != nil {
		return nil, fmt.Errorf("c.WaitForResponse() error: %w", err)
	}
	if len(resp.GetError()) > 0 {
		return nil, fmt.Errorf("sc2 client response error: %+v", resp.GetError())
	}
	return resp.GetCreateGame(), nil
}

func (c *RpcClient) JoinGame(ctx context.Context, req *sc2proto.RequestJoinGame) (*sc2proto.ResponseJoinGame, error) {
	id, err := c.SendRequest(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_JoinGame{
			JoinGame: req,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("c.SendRequest() error: %w", err)
	}
	resp, err := c.WaitForResponse(id)
	if err != nil {
		return nil, fmt.Errorf("c.WaitForResponse() error: %w", err)
	}
	if len(resp.GetError()) > 0 {
		return nil, fmt.Errorf("sc2 client response error: %+v", resp.GetError())
	}
	return resp.GetJoinGame(), nil
}

func (c *RpcClient) RestartGame(ctx context.Context, req *sc2proto.RequestRestartGame) (*sc2proto.ResponseRestartGame, error) {
	id, err := c.SendRequest(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_RestartGame{
			RestartGame: req,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("c.SendRequest() error: %w", err)
	}
	resp, err := c.WaitForResponse(id)
	if err != nil {
		return nil, fmt.Errorf("c.WaitForResponse() error: %w", err)
	}
	if len(resp.GetError()) > 0 {
		return nil, fmt.Errorf("sc2 client response error: %+v", resp.GetError())
	}
	return resp.GetRestartGame(), nil
}

func (c *RpcClient) LeaveGame(ctx context.Context, req *sc2proto.RequestLeaveGame) (*sc2proto.ResponseLeaveGame, error) {
	id, err := c.SendRequest(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_LeaveGame{
			LeaveGame: req,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("c.SendRequest() error: %w", err)
	}
	resp, err := c.WaitForResponse(id)
	if err != nil {
		return nil, fmt.Errorf("c.WaitForResponse() error: %w", err)
	}
	if len(resp.GetError()) > 0 {
		return nil, fmt.Errorf("sc2 client response error: %+v", resp.GetError())
	}
	return resp.GetLeaveGame(), nil
}

func (c *RpcClient) Quit(ctx context.Context, req *sc2proto.RequestQuit) (*sc2proto.ResponseQuit, error) {
	id, err := c.SendRequest(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_Quit{
			Quit: req,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("c.SendRequest() error: %w", err)
	}
	resp, err := c.WaitForResponse(id)
	if err != nil {
		return nil, fmt.Errorf("c.WaitForResponse() error: %w", err)
	}
	if len(resp.GetError()) > 0 {
		return nil, fmt.Errorf("sc2 client response error: %+v", resp.GetError())
	}
	return resp.GetQuit(), nil
}

func (c *RpcClient) Step(ctx context.Context, req *sc2proto.RequestStep) (*sc2proto.ResponseStep, error) {
	id, err := c.SendRequest(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_Step{
			Step: req,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("c.SendRequest() error: %w", err)
	}
	resp, err := c.WaitForResponse(id)
	if err != nil {
		return nil, fmt.Errorf("c.WaitForResponse() error: %w", err)
	}
	if len(resp.GetError()) > 0 {
		return nil, fmt.Errorf("sc2 client response error: %+v", resp.GetError())
	}
	return resp.GetStep(), nil
}

func (c *RpcClient) GameInfo(ctx context.Context, req *sc2proto.RequestGameInfo) (*sc2proto.ResponseGameInfo, error) {
	id, err := c.SendRequest(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_GameInfo{
			GameInfo: req,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("c.SendRequest() error: %w", err)
	}
	resp, err := c.WaitForResponse(id)
	if err != nil {
		return nil, fmt.Errorf("c.WaitForResponse() error: %w", err)
	}
	if len(resp.GetError()) > 0 {
		return nil, fmt.Errorf("sc2 client response error: %+v", resp.GetError())
	}
	return resp.GetGameInfo(), nil
}

func (c *RpcClient) Action(ctx context.Context, req *sc2proto.RequestAction) (*sc2proto.ResponseAction, error) {
	id, err := c.SendRequest(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_Action{
			Action: req,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("c.SendRequest() error: %w", err)
	}
	resp, err := c.WaitForResponse(id)
	if err != nil {
		return nil, fmt.Errorf("c.WaitForResponse() error: %w", err)
	}
	if len(resp.GetError()) > 0 {
		return nil, fmt.Errorf("sc2 client response error: %+v", resp.GetError())
	}
	return resp.GetAction(), nil
}

func (c *RpcClient) Observation(ctx context.Context, req *sc2proto.RequestObservation) (*sc2proto.ResponseObservation, error) {
	id, err := c.SendRequest(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_Observation{
			Observation: req,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("c.SendRequest() error: %w", err)
	}
	resp, err := c.WaitForResponse(id)
	if err != nil {
		return nil, fmt.Errorf("c.WaitForResponse() error: %w", err)
	}
	if len(resp.GetError()) > 0 {
		return nil, fmt.Errorf("sc2 client response error: %+v", resp.GetError())
	}
	return resp.GetObservation(), nil
}
