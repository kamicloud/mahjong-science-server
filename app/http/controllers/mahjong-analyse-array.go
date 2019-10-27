package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"mahjong-science-server/app/exceptions"
	"mahjong-science-server/app/http/dtos"
	"mahjong-science-server/app/http/services"
)

type MahjongAnalyseArrayController struct {
	beego.Controller
}

func (c *MahjongAnalyseArrayController) Post() {

	var request dtos.AnalyseArrayRequest
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

	message := dtos.AnalyseArrayMessage{
		Request: request,
	}

	exception := services.AnalyseArray(&message)

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
