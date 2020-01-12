package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/kamicloud/mahjong-science-server/app"
	_ "github.com/kamicloud/mahjong-science-server/app/console"
	"github.com/kamicloud/mahjong-science-server/routers"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now())
	logrus.Info("========== System Start ==========")
	e := echo.New()

	routers.Routes(e)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {
			if err := next(c); err != nil {
				return c.HTML(http.StatusNotFound, "备案中，网站关闭。")
			}

			return nil
		}
	})

	e.Logger.Fatal(e.Start(":8080"))
}
