package commands

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/api"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/proto/lq"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/tool"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"kamicloud/mahjong-science-server/app/utils"
	"os"
	"time"
)

const Players4 uint32 = 1
const Players3 uint32 = 2

func SyncRank() {
	fmt.Println("Every hour on the half hour")

	if err := syncRank(Players3); err != nil {
	}
	if err := syncRank(Players4); err != nil {
	}
}

type PlayerInfos []*lq.PlayerBaseView

func genReqLogin(username string, password string) (*lq.ReqLogin, error) {
	const key = "lailai" // from code.js
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(password))
	password = fmt.Sprintf("%x", mac.Sum(nil))

	// randomKey 最好是个固定值
	randomKey, ok := os.LookupEnv("RANDOM_KEY")
	if !ok {
		rawRandomKey, _ := uuid.NewV4()
		randomKey = rawRandomKey.String()
	}

	version, err := tool.GetMajsoulVersion(tool.ApiGetVersionZH)
	if err != nil {
		return nil, err
	}
	return &lq.ReqLogin{
		Account:   username,
		Password:  password,
		Reconnect: false,
		Device: &lq.ClientDeviceInfo{
			DeviceType: "pc",
			Os:         "",
			OsVersion:  "",
			Browser:    "safari",
		},
		RandomKey:         randomKey,          // 例如 aa566cfc-547e-4cc0-a36f-2ebe6269109b
		ClientVersion:     version.ResVersion, // 0.5.162.w
		GenAccessToken:    true,
		CurrencyPlatforms: []uint32{2}, // 1-inGooglePlay, 2-inChina
	}, nil
}

func syncRank(tp uint32) error {
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")

	endpoint, err := tool.GetMajsoulWebSocketURL()
	if err != nil {
		return err
	}
	c := api.NewWebSocketClient()
	if err := c.Connect(endpoint, tool.MajsoulOriginURL); err != nil {
		return err
	}
	defer c.Close()

	// 登录
	reqLogin, err := genReqLogin(username, password)
	if err != nil {
		return err
	}
	if _, err := c.Login(reqLogin); err != nil {
		return err
	}
	defer c.Logout(&lq.ReqLogout{})

	resp, err := c.FetchLevelLeaderboard(&lq.ReqLevelLeaderboard{
		Type: tp,
	})
	if err != nil {
		return err
	}

	var leaderPlayers = make(map[uint32]*lq.PlayerBaseView)

	for i := 0; i < len(resp.Items); i += 20 {
		items := resp.Items[i: i + 20]
		var ids []uint32
		for _, v := range items {
			ids = append(ids, v.GetAccountId())
		}
		players, err := fetchUserProfiles(c, ids)
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
	err = bm.Put(memKey, cacheValue, 86400 * 2 * time.Second)

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
