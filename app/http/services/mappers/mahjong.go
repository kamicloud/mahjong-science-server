package mappers

import (
	"github.com/EndlessCheng/mahjong-helper/util"
	"github.com/EndlessCheng/mahjong-helper/util/model"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"strings"
)

func MapAnalyseResult(
	playerInfo *model.PlayerInfo,
	shanten int,
	results util.Hand14AnalysisResultList,
	incShantenResults util.Hand14AnalysisResultList) dtos.TileAnalyseResult {

	var choices = resultToChoices(results)
	var incShantenChoices = resultToChoices(incShantenResults)

	res := dtos.TileAnalyseResult{
		CurrentTileString:       util.Tiles34ToStr(playerInfo.HandTiles34),
		CurrentTileSimpleString: Tiles34ToStr(playerInfo.HandTiles34),
		Shanten:                 shanten,
		CurrentTiles:            playerInfo.HandTiles34,
		CurrentRenderTiles:      util.Tiles34ToTiles(playerInfo.HandTiles34),
		Choices:                 choices,
		IncShantenChoices:       incShantenChoices,
	}

	return res
}

func resultToChoices(results []*util.Hand14AnalysisResult) (choices []dtos.DiscardChoice) {
	for _, result := range results {
		choice := *new(dtos.DiscardChoice)

		choice.Discard = result.DiscardTile
		choice.DrawCount, choice.Draws = result.Result13.Waits.ParseIndex()

		choices = append(choices, choice)
	}
	sortResult(choices)
	return
}

func sortResult(results []dtos.DiscardChoice) {
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[i].DrawCount < results[j].DrawCount {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
}

func Tiles34ToStr(tiles34 []int) (humanTiles string) {
	merge := func(lowerIndex, upperIndex int, endsWith string) {
		found := false
		for i, c := range tiles34[lowerIndex:upperIndex] {
			for j := 0; j < c; j++ {
				found = true
				humanTiles += string('1' + i)
			}
		}
		if found {
			humanTiles += endsWith
		}
	}
	merge(0, 9, "m")
	merge(9, 18, "p")
	merge(18, 27, "s")
	merge(27, 34, "z")
	return strings.TrimSpace(humanTiles)
}
