package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kamicloud/mahjong-science-server/app/http/controllers"
)

func registerApi(g *gin.Engine) {
	g.POST("/ping", controllers.Analyse)
}