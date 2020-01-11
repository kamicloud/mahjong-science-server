package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type GameLiveType struct {
	ID       int    `json:"id"`
	Name1Chs string `json:"name1_chs"`
	Name2Chs string `json:"name2_chs"`
}

type GameLiveFile struct {
	SelectFilters struct {
		Rows_ []GameLiveType `json:"rows_"`
	} `json:"select_filters"`
}

func GetGameLiveTypes() []GameLiveType {
	file, err := os.OpenFile("./conf/game_live.json", os.O_RDONLY, 0600)
	if err != nil {

	}
	data, err := ioutil.ReadAll(file)

	var gameLiveFile GameLiveFile
	//var gameLiveFile map[string]interface{}

	err = json.Unmarshal(data, &gameLiveFile)

	return gameLiveFile.SelectFilters.Rows_
}
