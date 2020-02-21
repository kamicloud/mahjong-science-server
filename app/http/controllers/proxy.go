package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
)

// Proxy 代理访问牌谱屋
func Proxy(c echo.Context) error {
	url := c.QueryParam("url")

	// url = strings.Replace(url, "ak-data-2.sapk.ch", "ak-data-1.sapk.ch", 1)

	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		if err != nil {
			logrus.Error(err)
			return c.JSON(200, dtos.BaseMessage{
				Status:  0,
				Message: "",
				Data:    nil,
			})
		}
	}
	client := &http.Client{}

	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		logrus.Error(err)
		return c.JSON(200, dtos.BaseMessage{
			Status:  0,
			Message: "",
			Data:    nil,
		})
	}
	var value []byte
	value, err = ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error(err)
		return c.JSON(200, dtos.BaseMessage{
			Status:  0,
			Message: "",
			Data:    nil,
		})
	}
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
