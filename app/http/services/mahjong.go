package services

import (
	"github.com/EndlessCheng/mahjong-helper/util"
	"github.com/EndlessCheng/mahjong-helper/util/model"
	"kamicloud/mahjong-science-server/app/exceptions"
	"kamicloud/mahjong-science-server/app/http/dtos"
	"kamicloud/mahjong-science-server/app/http/services/mappers"
	"kamicloud/mahjong-science-server/app/managers"
	"strconv"
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

	tileCount = util.CountOfTiles34(tiles34)

	if tileCount > 14 {
		return &exceptions.Exception{Status: exceptions.CustomError, Message: "输入手牌数量不符" + strconv.Itoa(tileCount)}
	}

	playerInfo := model.NewSimplePlayerInfo(tiles34, nil)

	util.CountOfTiles34(playerInfo.HandTiles34)
	// 分析手牌
	shanten, results14, incShantenResults := util.CalculateShantenWithImproves14(playerInfo)

	result := mappers.MapAnalyseResult(playerInfo, shanten, results14, incShantenResults)

	message.Response.Result = result

	return nil
}

func AnalyseArray(message *dtos.AnalyseArrayMessage) *exceptions.Exception {
	request := message.Request
	tiles34 := request.Tiles

	tileCount := util.CountOfTiles34(tiles34)

	if tileCount%3 == 0 || tileCount%3 == 1 {
		util.RandomAddTile(tiles34)
	}

	tileCount = util.CountOfTiles34(tiles34)

	if tileCount > 14 {
		return &exceptions.Exception{Status: exceptions.CustomError, Message: "输入手牌数量不符" + strconv.Itoa(tileCount)}
	}

	playerInfo := model.NewSimplePlayerInfo(tiles34, nil)

	util.CountOfTiles34(playerInfo.HandTiles34)
	// 分析手牌
	shanten, results14, incShantenResults := util.CalculateShantenWithImproves14(playerInfo)

	result := mappers.MapAnalyseResult(playerInfo, shanten, results14, incShantenResults)

	message.Response = dtos.AnalyseArrayResponse{
		Result: result,
	}

	return nil
}


func Random(message *dtos.RandomMessage) *exceptions.Exception {
	var tiles34 = managers.RandomTile34()

	var playerInfo *model.PlayerInfo

	playerInfo = model.NewSimplePlayerInfo(tiles34, nil)

	util.CountOfTiles34(playerInfo.HandTiles34)
	// 分析手牌
	shanten, results14, incShantenResults := util.CalculateShantenWithImproves14(playerInfo)

	message.Response = dtos.RandomResponse{
		Result: mappers.MapAnalyseResult(playerInfo, shanten, results14, incShantenResults),
	}
	return nil
}
