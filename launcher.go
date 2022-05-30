package sc2client

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Launcher struct {
	clientPath   string
	installDir   string
	workDir      string
	tempDir      string
	clientHost   string
	clientPort   int
	displayMode  int
	windowWidth  int
	windowHeight int
	windowX      int
	windowY      int
	cmd          *exec.Cmd
}

func NewLauncher(clientHost string, clientPort int, displayMode int, windowWidth int, windowHeight int, windowX int, windowY int) (*Launcher, error) {
	clientPath, err := GetSC2ClientPath()
	if err != nil {
		return nil, fmt.Errorf("GetSC2ClientPath() error: %w", err)
	}
	clientInstallDir, err := filepath.Abs(filepath.Join(clientPath, "..", "..", ".."))
	if err != nil {
		return nil, fmt.Errorf("filepath.Abs() error: %w", err)
	}
	workDir := filepath.Join(clientInstallDir, "Support")
	if strings.Contains(filepath.Base(clientPath), "x64") {
		workDir = filepath.Join(clientInstallDir, "Support64")
	}
	tempDir, err := os.MkdirTemp(os.TempDir(), "SC2_")
	if err != nil {
		return nil, fmt.Errorf("os.MkdirTemp() error: %w", err)
	}
	return &Launcher{
		clientPath:   clientPath,
		installDir:   clientInstallDir,
		workDir:      workDir,
		tempDir:      tempDir,
		clientHost:   clientHost,
		clientPort:   clientPort,
		displayMode:  displayMode,
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		windowX:      windowX,
		windowY:      windowY,
	}, nil
}

func (m *Launcher) StartProcess(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, m.clientPath,
		"-listen", m.clientHost,
		"-port", strconv.Itoa(m.clientPort),
		"-displayMode", strconv.Itoa(m.displayMode),
		"-windowwidth", strconv.Itoa(m.windowWidth),
		"-windowheight", strconv.Itoa(m.windowHeight),
		"-windowx", strconv.Itoa(m.windowX),
		"-windowy", strconv.Itoa(m.windowY),
		"-dataDir", m.installDir,
		"-tempDir", m.tempDir,
		"-verbose")
	cmd.Dir = m.workDir
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("cmd.Run() error: %w", err)
	}
	m.cmd = cmd
	return nil
}

func (m *Launcher) ProcessPid() int {
	return m.cmd.Process.Pid
}

func (m *Launcher) StopProcess() error {
	if err := m.cmd.Process.Kill(); err != nil {
		return fmt.Errorf("m.cmd.Process.Kill() error: %w", err)
	}
	return nil
}
