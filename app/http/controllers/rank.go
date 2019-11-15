package controllers

import (
	"encoding/json"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"github.com/kamicloud/mahjong-science-server/app/utils"
	"github.com/labstack/echo"
)

func Rank(c echo.Context) error {

	bm := utils.Cache

	rank4 := &[]*dtos.Rank{}
	rank3 := &[]*dtos.Rank{}

	cache4, found4 := bm.Get("rank4")
	cache3, found3 := bm.Get("rank3")

	if !found3 || !found4 {
		return c.JSON(200, dtos.BaseMessage{
			Status:  400,
			Message: "failed",
			Data:    nil,
		})
	}

	json.Unmarshal(cache4.([]byte), rank4)
	json.Unmarshal(cache3.([]byte), rank3)

	return c.JSON(200, dtos.BaseMessage{
		Status:  0,
		Message: "success",
		Data:    map[string]interface{}{
			"rank3": rank3,
			"rank4": rank4,
		},
	})
}
