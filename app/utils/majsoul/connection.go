package majsoul

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/api"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/proto/lq"
	"github.com/EndlessCheng/mahjong-helper/platform/majsoul/tool"
	"github.com/kamicloud/mahjong-science-server/app"
	"github.com/kamicloud/mahjong-science-server/app/exceptions"
	uuid "github.com/satori/go.uuid"
	"os"
	"sync"
)

var username string
var password string
var client *api.WebSocketClient

func init() {
	username = app.Config.Username
	password = app.Config.Password
}

func Login() error {
	reqLogin, err := genReqLogin(username, password)

	if err != nil {
		return err
	}

	_, err = client.Login(reqLogin)

	return err
}

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

var clinetMutex sync.Mutex

func GetClient(reconnect bool) (*api.WebSocketClient, error) {
	var err error
	clinetMutex.Lock()
	defer clinetMutex.Unlock()
	if reconnect && !CheckConnection() {
		client, err = getClient()

		if err != nil {
			return nil, err
		}
	}

	if client == nil {
		return nil, &exceptions.MajsoulConnectionError{}
	}

	return client, nil
}

func CheckConnection() bool {
	if client == nil {
		return false
	}

	_, err := client.Heatbeat(&lq.ReqHeatBeat{})

	return err == nil
}

func Close() error {
	if client == nil {
		return nil
	}

	err := client.Close()

	if err != nil {
		return err
	}
	client = nil

	return nil
}

func getClient() (*api.WebSocketClient, error) {
	endpoint, err := tool.GetMajsoulWebSocketURL()
	if err != nil {
		return nil, err
	}
	var client = api.NewWebSocketClient()
	if err := client.Connect(endpoint, tool.MajsoulOriginURL); err != nil {
		return nil, err
	}

	return client, nil
}
