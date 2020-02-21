package controllers

import (
	"context"
	"strconv"

	"github.com/kamicloud/mahjong-science-server/app/utils"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

// Stats 牌谱状态
func Stats(c echo.Context) error {
	gameLiveTypes := utils.GetGameLiveTypes()

	total := 0

	for _, gameLiveType := range *gameLiveTypes {
		roomID := strconv.Itoa(gameLiveType.ID)
		collection := utils.GetCollection("majsoul", "records_"+roomID)

		count, _ := collection.CountDocuments(context.TODO(), bson.M{})

		total += int(count)
	}

	return nil
}
