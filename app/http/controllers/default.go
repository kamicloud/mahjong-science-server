package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Default(c echo.Context) error {
	return c.HTML(http.StatusNotFound, "备案中，网站关闭。")
}
