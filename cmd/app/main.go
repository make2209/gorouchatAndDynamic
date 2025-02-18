package main

import (
	"github.com/gin-gonic/gin"
	"groupChatAndDynamic/api"
	"groupChatAndDynamic/inits"
)

func main() {
	r := gin.Default()
	inits.InitViper()
	inits.InitMysql()
	inits.InitRedis()
	api.LoadRouters(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
