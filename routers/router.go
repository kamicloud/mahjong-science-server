package routers

import (
	"github.com/kamicloud/mahjong-science-server/app/http/controllers"
	"github.com/labstack/echo"
)

func init() {
}

func Routes(e *echo.Echo) {
	e.GET("/", controllers.Default)
	e.POST("/mahjong/analyse", controllers.MahjongAnalyse)
	e.POST("/mahjong/analyse-array", controllers.MahjongAnalyseArray)
	e.POST("/mahjong/random", controllers.Random)
	e.POST("/wechat/code-to-session", controllers.WechatCodeToSession)
	e.GET("/mahjong/group", controllers.Group)
	e.GET("/mahjong/proxy", controllers.Proxy)
	e.GET("/mahjong/rank", controllers.Rank)
	e.GET("/records/stats", controllers.Stats)
}
