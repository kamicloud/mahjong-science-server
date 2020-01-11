package commands

import (
	"encoding/json"
	"fmt"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/api"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/proto/lq"
	"github.com/kamicloud/mahjong-science-server/app/utils"
	"github.com/kamicloud/mahjong-science-server/app/utils/majsoul"

	"time"
)

const Players4 uint32 = 1
const Players3 uint32 = 2

func SyncRank() {
	fmt.Println("Command SyncRank")

	if err := syncRank(Players3); err != nil {
		fmt.Println(err)
	}
	if err := syncRank(Players4); err != nil {
		fmt.Println(err)
	}
}

type PlayerInfos []*lq.PlayerBaseView

type GameLiveModel struct {
	lq.GameLiveHead
	_id string
}

func syncRank(tp uint32) error {
	client, err := majsoul.GetClient()

	if err != nil {
		return err
	}

	resp, err := client.FetchLevelLeaderboard(&lq.ReqLevelLeaderboard{
		Type: tp,
	})
	if err != nil {
		return err
	}

	var leaderPlayers = make(map[uint32]*lq.PlayerBaseView)

	for i := 0; i < len(resp.Items); i += 20 {
		items := resp.Items[i : i+20]
		var ids []uint32
		for _, v := range items {
			ids = append(ids, v.GetAccountId())
		}
		players, err := fetchUserProfiles(client, ids)
		if err != nil {
			return err
		}

		for _, v := range players {
			leaderPlayers[v.AccountId] = v
		}
	}

	var res []*lq.PlayerBaseView

	for _, v := range resp.Items {
		playerBaseView := leaderPlayers[v.AccountId]
		res = append(res, playerBaseView)
	}

	cacheValue, err := json.Marshal(res)

	if err != nil {
		return err
	}

	bm := utils.Cache
	var memKey string
	if tp == Players4 {
		memKey = "rank4"
	} else {
		memKey = "rank3"
	}
	bm.Set(memKey, cacheValue, 86400*2*time.Second)

	return nil
}

func fetchUserProfiles(c *api.WebSocketClient, ids []uint32) ([]*lq.PlayerBaseView, error) {
	resp, err := c.FetchMultiAccountBrief(&lq.ReqMultiAccountId{
		AccountIdList: ids,
	})

	if err != nil {
		return nil, err
	}

	return resp.Players, nil
}
