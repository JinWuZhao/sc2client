package sc2client

import (
	"context"
	"testing"
	"time"

	"github.com/JinWuZhao/sc2client/sc2proto"
)

func TestRpcClient_Ping(t *testing.T) {
	launcher, err := NewLauncher("127.0.0.1", 8167, 0, 1024, 768, 100, 100)
	if err != nil {
		t.Errorf("NewLauncher() error: %s", err)
		return
	}
	err = launcher.StartProcess(context.Background())
	if err != nil {
		t.Errorf("StartProcess() error: %s", err)
		return
	}
	t.Logf("launcher.StartProcess() return pid: %d", launcher.ProcessPid())

	time.Sleep(time.Second * 10)

	ctx := context.Background()
	conn, err := DialSC2(ctx, "127.0.0.1", 8167)
	if err != nil {
		t.Errorf("DialSC2() error: %s", err)
		return
	}
	rpcCli := NewRpcClient(conn, 10*time.Second)

	resp, err := rpcCli.Ping(ctx)
	if err != nil {
		t.Errorf("rpcCli.Ping() error: %s", err)
		return
	}
	t.Logf("rsp: %s, %d, %s, %d", resp.GetGameVersion(), resp.GetBaseBuild(), resp.GetDataVersion(), resp.GetDataBuild())

	_, err = rpcCli.Quit(ctx, &sc2proto.RequestQuit{})
	if err != nil {
		t.Errorf("rpcCli.Quit() error: %s", err)
		return
	}
	conn.Close()
	time.Sleep(time.Second * 10)
	err = launcher.StopProcess()
	if err != nil {
		t.Errorf("StopProcess() error: %s", err)
		return
	}
}
