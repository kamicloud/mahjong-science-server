package controllers

import (
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"github.com/labstack/echo"
)

func Default(c echo.Context) error {
	return c.JSON(200, dtos.BaseMessage{
		Status:  0,
		Message: "success",
		Data:    nil,
	})
}
