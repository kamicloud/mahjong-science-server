package main

import (
	"github.com/astaxie/beego"
	_ "mahjong-science-server/app"
	_ "mahjong-science-server/app/console"
	_ "mahjong-science-server/routers"
)

func main() {
	beego.Run()
}

