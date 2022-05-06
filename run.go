package sc2client

import (
	"context"
	"fmt"
	"sync"
)

func RunGame(ctx context.Context, gameMap string, players []*PlayerSetup, disableFog bool) error {
	if gameMap == "" {
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
		gameLoop:
			for {
				if index == 0 {
					err = client.HostGame(ctx, pc, gameMap, players, disableFog)
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
