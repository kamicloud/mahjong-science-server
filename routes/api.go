package routes

import (
	"awesomeProject/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func registerApi(g *gin.Engine) {
	g.POST("/ping", controllers.Analyse)
}