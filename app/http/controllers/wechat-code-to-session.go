package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/miniprogram"
	"mahjong-science-server/app/exceptions"
	"mahjong-science-server/app/http/dtos"
)

var wxa *miniprogram.MiniProgram

func init() {
	memCache := cache.NewMemcache(beego.AppConfig.String("memcached"))
	config := &wechat.Config{
		AppID:     beego.AppConfig.String("appId"),
		AppSecret: beego.AppConfig.String("appSecret"),
		Cache:     memCache,
	}
	wc := wechat.NewWechat(config)

	wxa = wc.GetMiniProgram()
}


type WechatCodeToSessionController struct {
	beego.Controller
}

func (c *WechatCodeToSessionController) Post() {
	var request dtos.CodeToSessionRequest
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

	result, err := wxa.Code2Session(request.Code)

	c.Data["json"] = result

	c.ServeJSON()
}
