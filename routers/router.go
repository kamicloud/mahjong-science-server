package routers

import (
	"github.com/astaxie/beego"
	controllers2 "mahjong-science-server/app/http/controllers"
)

func init() {
    beego.Router("/", &controllers2.MainController{})

	beego.Router("/mahjong/random", &controllers2.RandomController{})
	beego.Router("/mahjong/analyse", &controllers2.MahjongAnalyseController{})
	beego.Router("/mahjong/analyse-array", &controllers2.MahjongAnalyseArrayController{})
}
