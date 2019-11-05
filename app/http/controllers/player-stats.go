package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

type PlayerStatsController struct {
	beego.Controller
}

func (c *PlayerStatsController) Get() {
	client := &http.Client{}
	id := c.Ctx.Input.Param(":id")

	//生成要访问的url
	url := "https://ak-data-2.sapk.ch/api/player_stats/" + id + "?mode="

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)
	var value []byte
	value, _ = ioutil.ReadAll(response.Body)
	jsonobj := map[string]interface{}{}
	json.Unmarshal(value, &jsonobj)
	c.Data["json"] = jsonobj
	c.ServeJSON()
}
