package sc2client

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/jinwuzhao/sc2client/sc2proto"
)

func TestRun_RunGame(t *testing.T) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	err := RunGame(ctx,
		"Custom/WarBattle.SC2Map",
		[]*PlayerSetup{
			{
				Type:       sc2proto.PlayerType_Participant,
				Race:       sc2proto.Race_Random,
				Name:       "Director",
				Difficulty: sc2proto.Difficulty_Easy,
				AIBuild:    sc2proto.AIBuild_RandomBuild,
				Step: func(ctx context.Context, steps uint32, rpc *RpcClient, stop chan<- error) {
					if steps%100 == 0 {
						_, err := rpc.Action(ctx, &sc2proto.RequestAction{
							Actions: []*sc2proto.Action{
								{
									ActionChat: &sc2proto.ActionChat{
										Channel: sc2proto.ActionChat_Team.Enum(),
										Message: proto.String(fmt.Sprintf("cmd-create-siege-tank 3 红-%d", steps)),
									},
								},
								{
									ActionChat: &sc2proto.ActionChat{
										Channel: sc2proto.ActionChat_Team.Enum(),
										Message: proto.String(fmt.Sprintf("cmd-move-toward 红-%d %d %d", steps, rand.Intn(90)-45, rand.Intn(100)*2)),
									},
								},
								{
									ActionChat: &sc2proto.ActionChat{
										Channel: sc2proto.ActionChat_Team.Enum(),
										Message: proto.String(fmt.Sprintf("cmd-create-siege-tank 4 蓝-%d", steps)),
									},
								},
								{
									ActionChat: &sc2proto.ActionChat{
										Channel: sc2proto.ActionChat_Team.Enum(),
										Message: proto.String(fmt.Sprintf("cmd-move-toward 蓝-%d %d %d", steps, rand.Intn(90)+135, rand.Intn(100)*2)),
									},
								},
							},
						})
						if err != nil {
							stop <- fmt.Errorf("rpc.Action() error: %w", err)
							return
						}
					}
				},
			},
			{
				Type:       sc2proto.PlayerType_Participant,
				Race:       sc2proto.Race_Random,
				Name:       "Audience",
				Difficulty: sc2proto.Difficulty_Easy,
				AIBuild:    sc2proto.AIBuild_RandomBuild,
			},
		},
		true)
	if err != nil {
		t.Error(err)
	}
}
