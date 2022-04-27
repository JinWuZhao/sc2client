package sc2client

import (
	"fmt"
	"net"
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
