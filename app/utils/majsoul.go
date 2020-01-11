package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// GameLiveType 观战类型
type GameLiveType struct {
	ID       int    `json:"id"`
	Name1Chs string `json:"name1_chs"`
	Name2Chs string `json:"name2_chs"`
}

// GameLiveFile 观战文件
type GameLiveFile struct {
	SelectFilters struct {
		Rows []GameLiveType `json:"rows_"`
	} `json:"select_filters"`
}

var gameLiveTypes *[]GameLiveType

// GetGameLiveTypes 获取游戏房间类型
func GetGameLiveTypes() *[]GameLiveType {
	if gameLiveTypes == nil {
		file, err := os.OpenFile("./conf/game_live.json", os.O_RDONLY, 0600)
		if err != nil {

		}
		data, err := ioutil.ReadAll(file)

		var gameLiveFile GameLiveFile
		//var gameLiveFile map[string]interface{}

		err = json.Unmarshal(data, &gameLiveFile)

		gameLiveTypes = &gameLiveFile.SelectFilters.Rows
	}

	return gameLiveTypes
}
