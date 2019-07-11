package main

import (
	"github.com/gin-gonic/gin"
	"mahjong-science-server/routes"
)

func main() {
	r := gin.Default()
	routes.Register(r)
	r.Run() // 在 0.0.0.0:8080 上监听并服务
}