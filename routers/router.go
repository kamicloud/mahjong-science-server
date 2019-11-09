package routers

import (
	"github.com/astaxie/beego"
	"github.com/kamicloud/mahjong-science-server/app/http/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/mahjong/search-player/:name", &controllers.SearchPlayerController{})
	beego.Router("/mahjong/player-extended-stats/:id", &controllers.PlayerExtendedStatsController{})
	beego.Router("/mahjong/player-stats/:id", &controllers.PlayerStatsController{})

	beego.Router("/mahjong/random", &controllers.RandomController{})
	beego.Router("/mahjong/proxy", &controllers.ProxyController{})
	beego.Router("/mahjong/rank", &controllers.RankController{})
	beego.Router("/mahjong/group", &controllers.GroupController{})
	beego.Router("/mahjong/analyse", &controllers.MahjongAnalyseController{})
	beego.Router("/mahjong/analyse-array", &controllers.MahjongAnalyseArrayController{})
	beego.Router("/wechat/code-to-session", &controllers.WechatCodeToSessionController{})
}
