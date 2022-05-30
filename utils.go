package sc2client

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

func GetLocalAddress() (string, int, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", 0, fmt.Errorf("net.Listen error: %w", err)
	}
	defer listener.Close()
	host, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		return "", 0, fmt.Errorf("net.SplitHostPort error: %w", err)
	}
	portNum, _ := strconv.Atoi(port)
	return host, portNum, nil
}

type PortConfig struct {
	Servers [2]int
	Players [2]int
}

func NewPortConfig() (*PortConfig, error) {
	pc := new(PortConfig)
	var err error
	for i := range pc.Servers {
		_, pc.Servers[i], err = GetLocalAddress()
		if err != nil {
			return nil, fmt.Errorf("GetLocalAddress() error: %w", err)
		}
	}
	for i := range pc.Players {
		_, pc.Players[i], err = GetLocalAddress()
		if err != nil {
			return nil, fmt.Errorf("GetLocalAddress() error: %w", err)
		}
	}
	return pc, nil
}

func GetSC2ClientPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("os.UserHomeDir() error: %w", err)
	}
	exeInfo, err := os.ReadFile(filepath.Join(homeDir, "Documents", "StarCraft II", "ExecuteInfo.txt"))
	if err != nil {
		return "", fmt.Errorf("os.ReadFile() error: %w", err)
	}
	stmtParts := bytes.Split(exeInfo, []byte{'='})
	if len(stmtParts) != 2 && !bytes.Equal(bytes.TrimSpace(stmtParts[0]), []byte("executable")) {
		return "", fmt.Errorf("unrecognized ExecuteInfo: %s", exeInfo)
	}
	return string(bytes.TrimSuffix(bytes.TrimSpace(stmtParts[1]), []byte{'\n', 0x00})), nil
}

func GetSC2InstallDir() (string, error) {
	clientPath, err := GetSC2ClientPath()
	if err != nil {
		return "", fmt.Errorf("GetSC2ClientPath() error: %w", err)
	}
	clientInstallDir, err := filepath.Abs(filepath.Join(clientPath, "..", "..", ".."))
	if err != nil {
		return "", fmt.Errorf("filepath.Abs() error: %w", err)
	}
	return clientInstallDir, nil
}
