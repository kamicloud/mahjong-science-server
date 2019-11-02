package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
)

type GroupController struct {
	beego.Controller
}

func (c *GroupController) Post() {
	c.Data["json"] = dtos.BaseMessage{
		Status:  0,
		Message: "success",
		Data: dtos.GroupResponse{
			Groups: []dtos.Group{
				{Title: "日麻杂谈", Num: "375 865 038", Content: "日本麻将技术交流，牌谱探讨"},
			},
		},
	}
	c.ServeJSON()
}
