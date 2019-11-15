package main

import (
	_ "github.com/kamicloud/mahjong-science-server/app"
	_ "github.com/kamicloud/mahjong-science-server/app/console"
	"github.com/kamicloud/mahjong-science-server/routers"
	"github.com/labstack/echo"
)


func main() {
	e := echo.New()

	routers.Routes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

