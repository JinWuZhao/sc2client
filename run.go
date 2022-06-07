package sc2client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type GameMap struct {
	Name       string
	SourcePath string
}

func installGameMap(gameMap GameMap) error {
	if gameMap.SourcePath == "" {
		return nil
	}
	sc2Path, err := GetSC2InstallDir()
	if err != nil {
		return fmt.Errorf("GetSC2InstallDir() error: %w", err)
	}
	sc2MapPath := filepath.Join(sc2Path, "Maps")
	if _, err := os.Stat(sc2MapPath); os.IsNotExist(err) {
		err = os.Mkdir(sc2MapPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("os.Mkdir(%s) error: %w", sc2MapPath, err)
		}
	}
	content, err := os.ReadFile(gameMap.SourcePath)
	if err != nil {
		return fmt.Errorf("os.ReadFile(%s) error: %w", gameMap.SourcePath, err)
	}
	dstPath := filepath.Join(sc2MapPath, gameMap.Name)
	err = os.WriteFile(dstPath, content, os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.WriteFile(%s) error: %w", dstPath, err)
	}
	return nil
}

func RunGame(ctx context.Context, gameMaps []GameMap, players []*PlayerSetup, disableFog bool) error {
	if len(gameMaps) <= 0 {
		return fmt.Errorf("invalid game map")
	}
	if len(players) != 2 {
		return fmt.Errorf("need two players")
	}

	pc, err := NewPortConfig()
	if err != nil {
		return fmt.Errorf("NewPortConfig() error: %w", err)
	}

	clients := []*Client{NewClient(), NewClient()}
	errors := make([]error, 2)
	var wg sync.WaitGroup
	wg.Add(len(clients))
	for i, c := range clients {
		index := i
		client := c
		go func() {
			err := client.Init(ctx)
			if err != nil {
				errors[index] = fmt.Errorf("client.Init() error: %w", err)
				wg.Done()
				return
			}
			var mapIndex int
		gameLoop:
			for {
				if index == 0 {
					err = installGameMap(gameMaps[mapIndex])
					if err != nil {
						errors[index] = fmt.Errorf("installGameMap() error: %w", err)
						break gameLoop
					}
					err = client.HostGame(ctx, pc, gameMaps[mapIndex].Name, players, disableFog)
					if err != nil {
						errors[index] = fmt.Errorf("client.HostGame() error: %w", err)
						break gameLoop
					}
				} else {
					err = client.JoinGame(ctx, pc, players)
					if err != nil {
						errors[index] = fmt.Errorf("client.JoinGame() error: %w", err)
						break gameLoop
					}
				}
				client.StartGameLoop(ctx)
				err = client.WaitGameEnd()
				if err != nil {
					errors[index] = fmt.Errorf("client.WaitGameEnd() error: %w", err)
					break gameLoop
				}
				select {
				case <-ctx.Done():
					break gameLoop
				default:
				}
				mapIndex++
				if mapIndex >= len(gameMaps) {
					mapIndex = 0
				}
			}
			client.Finalize()
			wg.Done()
		}()
	}
	wg.Wait()

	if errors[0] != nil || errors[1] != nil {
		return fmt.Errorf("clients error: [1](%s), [2](%s)", errors[0], errors[1])
	}
	return nil
}
