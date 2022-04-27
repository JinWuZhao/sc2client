package sc2client

import (
	"context"
	"testing"
)

func TestLauncher_StartProcess(t *testing.T) {
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
}
