package main

import (
	"github.com/astaxie/beego"
	_ "kamicloud/mahjong-science-server/app"
	_ "kamicloud/mahjong-science-server/app/console"
	_ "kamicloud/mahjong-science-server/routers"
)

func main() {
	beego.Run()
}

