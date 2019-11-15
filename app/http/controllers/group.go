package controllers

import (
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"github.com/labstack/echo"
)

func Group(c echo.Context) error {
	return c.JSON(200, dtos.BaseMessage{
		Status:  0,
		Message: "success",
		Data: dtos.GroupResponse{
			Groups: []dtos.Group{
				{Title: "日麻杂谈", Num: "375 865 038", Content: "日本麻将技术交流，牌谱探讨"},
			},
		},
	})
}