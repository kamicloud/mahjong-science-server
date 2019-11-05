package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

type SearchPlayerController struct {
	beego.Controller
}

func (c *SearchPlayerController) Get() {
	client := &http.Client{}
	name := c.Ctx.Input.Param(":name")

	//生成要访问的url
	url := "https://ak-data-2.sapk.ch/api/search_player/" + name + "?limit=20"

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)
	var value []byte
	value, _ = ioutil.ReadAll(response.Body)
	jsonobj := []map[string]interface{}{}
	json.Unmarshal(value, &jsonobj)
	c.Data["json"] = jsonobj
	c.ServeJSON()
}
