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
	if len(players) < 2 {
		return fmt.Errorf("need two players at least")
	}

	pc, err := NewPortConfig()
	if err != nil {
		return fmt.Errorf("NewPortConfig() error: %w", err)
	}

	clients := []*Client{NewClient(), NewClient()}
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(len(clients))
	for i, c := range clients {
		index := i
		client := c
		go func() {
			err := client.Init(ctx)
			if err != nil {
				errors <- fmt.Errorf("client.Init() error: %w", err)
				return
			}
			for ctx.Err() == nil {
				if index == 0 {
					err = client.HostGame(ctx, pc, gameMap, players, disableFog)
					if err != nil {
						errors <- fmt.Errorf("client.HostGame() error: %w", err)
						return
					}
				} else {
					err = client.JoinGame(ctx, pc, players)
					if err != nil {
						errors <- fmt.Errorf("client.JoinGame() error: %w", err)
						return
					}
				}
				client.StartGameLoop(ctx)
				err = client.WaitGameEnd()
				if err != nil {
					errors <- fmt.Errorf("client.WaitGameEnd() error: %w", err)
					return
				}
			}
			client.Finalize()
			select {
			case errors <- nil:
			default:
			}
			wg.Done()
		}()
	}
	wg.Wait()

	err1 := <-errors
	err2 := <-errors
	if err1 != nil || err2 != nil {
		return fmt.Errorf("clients error: [1](%s), [2](%s)", err1, err2)
	}
	return nil
}
