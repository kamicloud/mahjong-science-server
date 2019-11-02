package routers

import (
	"github.com/astaxie/beego"
	"github.com/kamicloud/mahjong-science-server/app/http/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/mahjong/random", &controllers.RandomController{})
	beego.Router("/mahjong/rank", &controllers.RankController{})
	beego.Router("/mahjong/group", &controllers.GroupController{})
	beego.Router("/mahjong/analyse", &controllers.MahjongAnalyseController{})
	beego.Router("/mahjong/analyse-array", &controllers.MahjongAnalyseArrayController{})
	beego.Router("/wechat/code-to-session", &controllers.WechatCodeToSessionController{})
}
