package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"github.com/kamicloud/mahjong-science-server/app/utils"
)

type RankController struct {
	beego.Controller
}

func (c *RankController) Get() {
	bm := utils.Cache

	rank4 := &[]*dtos.Rank{}
	rank3 := &[]*dtos.Rank{}

	if !bm.IsExist("rank3") || !bm.IsExist("rank4") {
		c.Data["json"] = dtos.BaseMessage{
			Status:  400,
			Message: "failed",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	json.Unmarshal(bm.Get("rank4").([]byte), rank4)
	json.Unmarshal(bm.Get("rank4").([]byte), rank3)

	c.Data["json"] = dtos.BaseMessage{
		Status:  0,
		Message: "success",
		Data:    map[string]interface{}{
			"rank3": rank3,
			"rank4": rank4,
		},
	}
	c.ServeJSON()
}
