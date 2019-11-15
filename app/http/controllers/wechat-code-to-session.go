package controllers

import (
	"github.com/kamicloud/mahjong-science-server/app"
	"github.com/kamicloud/mahjong-science-server/app/exceptions"
	"github.com/kamicloud/mahjong-science-server/app/http/dtos"
	"github.com/labstack/echo"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/miniprogram"
)

var wxa *miniprogram.MiniProgram

func init() {
	memCache := cache.NewMemcache(app.Config.Memcached)
	config := &wechat.Config{
		AppID:     app.Config.AppId,
		AppSecret: app.Config.AppSecret,
		Cache:     memCache,
	}
	wc := wechat.NewWechat(config)

	wxa = wc.GetMiniProgram()
}

func WechatCodeToSession(c echo.Context) error {
	var request dtos.CodeToSessionRequest
	var err error

	if err = c.Bind(request); err != nil {
		// err
		return c.JSON(200, exceptions.Exception{
			Status:  exceptions.InvalidParameter,
			Message: "参数错误",
		})
	}

	result, err := wxa.Code2Session(request.Code)

	return c.JSON(200, result)
}

