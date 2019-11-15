package controllers

import (
	"github.com/kamicloud/mahjong-science-server/app/exceptions"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"github.com/kamicloud/mahjong-science-server/app/http/services"
	"github.com/labstack/echo"
)

func MahjongAnalyseArray(c echo.Context) error {

	var request = new(dtos.AnalyseArrayRequest)
	var err error

	if err = c.Bind(request); err != nil {
		// err
		return c.JSON(200, exceptions.Exception{
			Status:  exceptions.InvalidParameter,
			Message: "参数错误",
		})
	}

	message := dtos.AnalyseArrayMessage{
		Request: request,
	}

	exception := services.AnalyseArray(&message)

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
