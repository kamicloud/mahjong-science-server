package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/kamicloud/mahjong-science-server/app/exceptions"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"github.com/kamicloud/mahjong-science-server/app/http/services"
)

type RandomController struct {
	beego.Controller
}

func (c *RandomController) Post() {

	var request dtos.RandomRequest
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &request)

	if err != nil {
		// err
		c.Data["json"] = exceptions.Exception{
			Status:  exceptions.InvalidParameter,
			Message: "参数错误",
		}
		c.ServeJSON()
		return
	}

	message := dtos.RandomMessage{
		Request: request,
	}

	exception := services.Random(&message)

	if exception != nil {
		// err
		c.Data["json"] = exception
		c.ServeJSON()
		return
	}

	c.Data["json"] = dtos.BaseMessage{
		Status:  0,
		Message: "success",
		Data:    message.Response,
	}

	c.ServeJSON()
}