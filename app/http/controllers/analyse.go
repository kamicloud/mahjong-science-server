package controllers

import (
	"fmt"
	"github.com/EndlessCheng/mahjong-helper/util"
	"github.com/EndlessCheng/mahjong-helper/util/model"
	"github.com/gin-gonic/gin"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"strconv"
)


type FormPing struct {
	Tiles string `json:"tiles" binding:"required"`
}


func buildResponse(playerInfo *model.PlayerInfo, shanten int, results util.Hand14AnalysisResultList, incShantenResults util.Hand14AnalysisResultList) gin.H {
	var choices = make([]dtos.Choice, 0)
	var incShantenChoices = make([]dtos.Choice, 0)

	for _, result := range results {
		choice := *new(dtos.Choice)

		choice.Discard = result.DiscardTile
		choice.DrawCount, choice.Draws = result.Result13.Waits.ParseIndex()

		choices = append(choices, choice)
	}

	for _, incShantenResult := range incShantenResults {
		incShantenChoice := *new(dtos.Choice)

		incShantenChoice.Discard = incShantenResult.DiscardTile
		incShantenChoice.DrawCount, incShantenChoice.Draws = incShantenResult.Result13.Waits.ParseIndex()

		incShantenChoices = append(incShantenChoices, incShantenChoice)
	}

	return gin.H{
		"status": 0,
		"data": dtos.AnalyseResponse{
			Shanten:           shanten,
			CurrentTiles:      playerInfo.HandTiles34,
			Choices:           choices,
			IncShantenChoices: incShantenChoices,
		},
	}
}

func Analyse(c *gin.Context) {
	//param, success := c.GetPostForm("tiles")
	var param FormPing
	err := c.ShouldBindJSON(&param)

	if err != nil {
		c.JSON(200, gin.H{
			"status":  0,
			"message": "failed",
		})
		return
	}

	var playerInfo *model.PlayerInfo

	//tiles34, _, _ := util.StrToTiles34("244078m137p66789s")
	tiles34, _, err := util.StrToTiles34(param.Tiles)

	tileCount := util.CountOfTiles34(tiles34)

	if tileCount%3 == 0 {
		c.JSON(200, gin.H{
			"status":  1,
			"message": "输入牌数为" + strconv.Itoa(tileCount) + "，不合法",
		})
	}

	playerInfo = model.NewSimplePlayerInfo(tiles34, nil)

	util.CountOfTiles34(playerInfo.HandTiles34)
	// 分析手牌
	shanten, results14, incShantenResults := util.CalculateShantenWithImproves14(playerInfo)

	c.JSON(200, buildResponse(playerInfo, shanten, results14, incShantenResults))
}

func sortResults(results util.Hand14AnalysisResultList)  {
	for i:= 0; i < len(results);i++ {
		for j:= 0; j < len(results) ;j ++ {
			if results[i].Result13.AvgNextShantenWaitsCount < results[j].Result13.AvgNextShantenWaitsCount {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
}

func printResults14WithRisk(results14 util.Hand14AnalysisResultList) {
	if len(results14) == 0 {
		return
	}
	// FIXME: 选择很多时如何精简何切选项？
	const maxShown = 10
	shownResults14 := results14
	if len(shownResults14) > maxShown { // 限制输出数量
		shownResults14 = shownResults14[:maxShown]
	}
	if len(results14[0].OpenTiles) > 0 {
		fmt.Print("鸣牌后")
	}
	fmt.Println(util.NumberToChineseShanten(results14[0].Result13.Shanten) + "：")
	//for _, result := range results14 {
	// printWaitsWithImproves13_oneRow(result.Result13, result.DiscardTile, result.OpenTiles, mixedRiskTable)
	//}
}
