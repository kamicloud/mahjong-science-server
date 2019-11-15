package controllers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"

	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
)

func Proxy(c echo.Context) error {
	url := c.QueryParam("url")

	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}
	client := &http.Client{}

	//处理返回结果
	response, _ := client.Do(request)
	var value []byte
	value, _ = ioutil.ReadAll(response.Body)
	jsonobj := map[string]interface{}{}
	jsonarr := []map[string]interface{}{}
	json.Unmarshal(value, &jsonobj)
	json.Unmarshal(value, &jsonarr)
	if len(jsonobj) == 0 {
		return c.JSON(200, dtos.BaseMessage{
			Status:  0,
			Message: "",
			Data:    jsonarr,
		})
	} else {
		return c.JSON(200, dtos.BaseMessage{
			Status:  0,
			Message: "",
			Data:    jsonobj,
		})
	}
}

