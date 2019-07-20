package controllers

import (
	"github.com/EndlessCheng/mahjong-helper/util"
	"github.com/EndlessCheng/mahjong-helper/util/model"
	"github.com/gin-gonic/gin"
	"github.com/kamicloud/mahjong-science-server/app/exceptions"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"github.com/kamicloud/mahjong-science-server/app/http/mapper"
	"github.com/kamicloud/mahjong-science-server/app/http/services"
	"github.com/kamicloud/mahjong-science-server/app/manager"
)

func Analyse(c *gin.Context) {
	var request dtos.AnalyseRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(200, buildResponse(exceptions.InvalidParameter, "参数错误", nil))
		return
	}

	message := dtos.AnalyseMessage{
		Request: request,
	}

	exception := services.Analyse(&message)

	if exception != nil {
		c.JSON(200, buildResponse(exception.Status, exception.Message, nil))
		return
	}

	c.JSON(200, buildResponse(exceptions.Success, "success", message.Response))
}

func AnalyseArray(c *gin.Context) {
	var request dtos.AnalyseArrayRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(200, buildResponse(exceptions.InvalidParameter, "参数错误", nil))
		return
	}

	//message := dtos.AnalyseMessage{
	//	Request: request,
	//}

	response, exception := services.AnalyseArray(request)

	if exception != nil {
		c.JSON(200, buildResponse(exception.Status, exception.Message, nil))
		return
	}

	c.JSON(200, buildResponse(exceptions.Success, "success", response))
}

func Random(c *gin.Context) {
	var tiles34 = manager.RandomTile34()

	var playerInfo *model.PlayerInfo

	playerInfo = model.NewSimplePlayerInfo(tiles34, nil)

	util.CountOfTiles34(playerInfo.HandTiles34)
	// 分析手牌
	shanten, results14, incShantenResults := util.CalculateShantenWithImproves14(playerInfo)

	c.JSON(200, buildResponse(exceptions.Success, "success", dtos.AnalyseResponse{
		Result: mapper.MapAnalyseResult(playerInfo, shanten, results14, incShantenResults),
	}))
}

func Group(c *gin.Context) {
	c.JSON(200, buildResponse(exceptions.Success, "success", dtos.GroupResponse{
		Groups: []dtos.Group{
			{Title: "日麻杂谈", Num: "375 865 038", Content: "日本麻将技术交流，牌谱探讨"},
		},
	}))
}

func buildResponse(status int, message string, data interface{}) gin.H {
	return gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	}
}
