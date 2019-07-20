package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kamicloud/mahjong-science-server/app/http/controllers"
	"github.com/kamicloud/mahjong-science-server/app/http/middleware"
)

func registerApi(g *gin.Engine) {
	g1 := g.Group("/mahjong", middleware.TemplateMiddleware())
	{
		g1.POST("/ping", controllers.Analyse)
		g1.POST("/analyse", controllers.Analyse)
		g1.POST("/analyse-array", controllers.AnalyseArray)
		g1.POST("random", controllers.Random)
		g1.POST("group", controllers.Group)
	}
}
