package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
)

type ProxyController struct {
	beego.Controller
}

func (c *ProxyController) Get() {
	url := c.GetString("url")

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}
	client := &http.Client{}

	//处理返回结果
	response, _ := client.Do(reqest)
	var value []byte
	value, _ = ioutil.ReadAll(response.Body)
	jsonobj := map[string]interface{}{}
	jsonarr := []map[string]interface{}{}
	json.Unmarshal(value, &jsonobj)
	json.Unmarshal(value, &jsonarr)
	if len(jsonobj) == 0 {
		c.Data["json"] = dtos.BaseMessage{
			Status:  0,
			Message: "",
			Data:    jsonarr,
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = dtos.BaseMessage{
			Status:  0,
			Message: "",
			Data:    jsonobj,
		}
		c.ServeJSON()
	}
}
