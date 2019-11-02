package main

import (
	"github.com/astaxie/beego"
	_ "github.com/kamicloud/mahjong-science-server/app"
	_ "github.com/kamicloud/mahjong-science-server/app/console"
	_ "github.com/kamicloud/mahjong-science-server/routers"
)

func main() {
	beego.Run()
}

