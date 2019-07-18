package services

import (
	"github.com/EndlessCheng/mahjong-helper/util"
	"github.com/EndlessCheng/mahjong-helper/util/model"
	"github.com/kamicloud/mahjong-science-server/app/exceptions"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"github.com/kamicloud/mahjong-science-server/app/http/mapper"
)

func Analyse(message *dtos.AnalyseMessage) *exceptions.Exception {
	tiles := message.Request.Tiles

	tiles34, _, err := util.StrToTiles34(tiles)

	if tiles34 == nil || err != nil {
		return &exceptions.Exception{Status: exceptions.CustomError, Message: "输入格式非法"}
	}

	tileCount := util.CountOfTiles34(tiles34)

	if tileCount%3 == 0 || tileCount%3 == 1 {
		util.RandomAddTile(tiles34)
	}

	playerInfo := model.NewSimplePlayerInfo(tiles34, nil)

	util.CountOfTiles34(playerInfo.HandTiles34)
	// 分析手牌
	shanten, results14, incShantenResults := util.CalculateShantenWithImproves14(playerInfo)

	result := mapper.MapAnalyseResult(playerInfo, shanten, results14, incShantenResults)

	message.Response.Result = result

	return nil
}

func AnalyseArray(request dtos.AnalyseArrayRequest) (*dtos.AnalyseArrayResponse, *exceptions.Exception) {
	tiles34 := request.Tiles

	tileCount := util.CountOfTiles34(tiles34)

	if tileCount%3 == 0 || tileCount%3 == 1 {
		util.RandomAddTile(tiles34)
	}

	playerInfo := model.NewSimplePlayerInfo(tiles34, nil)

	util.CountOfTiles34(playerInfo.HandTiles34)
	// 分析手牌
	shanten, results14, incShantenResults := util.CalculateShantenWithImproves14(playerInfo)

	result := mapper.MapAnalyseResult(playerInfo, shanten, results14, incShantenResults)

	response := dtos.AnalyseArrayResponse{
		Result: result,
	}

	return &response, nil
}
