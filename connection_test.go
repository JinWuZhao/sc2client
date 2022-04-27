package sc2client

import (
	"context"
	"testing"
	"time"

	"github.com/jinwuzhao/sc2client/sc2proto"
)

func TestConnection_WriteRead(t *testing.T) {
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
	c, err := DialSC2(ctx, "127.0.0.1", 8167)
	if err != nil {
		t.Errorf("DialSC2() error: %s", err)
		return
	}
	err = c.Write(ctx, &sc2proto.Request{
		Request: &sc2proto.Request_Ping{
			Ping: &sc2proto.RequestPing{},
		},
	})
	if err != nil {
		t.Errorf("c.Write() error: %s", err)
		return
	}
	rsp := &sc2proto.Response{}
	err = c.Read(ctx, rsp)
	if err != nil {
		t.Errorf("c.Read() error: %s", err)
		return
	}
	rspPing := rsp.GetPing()
	t.Logf("rsp: %s, %d, %s, %d", rspPing.GetGameVersion(), rspPing.GetBaseBuild(), rspPing.GetDataVersion(), rspPing.GetDataBuild())

	c.Close()
	err = launcher.StopProcess()
	if err != nil {
		t.Errorf("StopProcess() error: %s", err)
		return
	}
}
