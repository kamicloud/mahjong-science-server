package controllers

import (
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"github.com/kamicloud/mahjong-science-server/app/http/services"
	"github.com/labstack/echo"
)

func Random(c echo.Context) error {
	var request dtos.RandomRequest

	message := dtos.RandomMessage{
		Request: request,
	}

	exception := services.Random(&message)

	if exception != nil {
		// err
		return c.JSON(200, exception)
	}

	return c.JSON(200, dtos.BaseMessage{
		Status:  0,
		Message: "success",
		Data:    message.Response,
	})
}

