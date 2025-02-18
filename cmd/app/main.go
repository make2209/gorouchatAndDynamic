package main

import (
	"github.com/gin-gonic/gin"
	"zk0212/api"
	"zk0212/inits"
)

func main() {
	r := gin.Default()
	inits.InitViper()
	inits.InitMysql()
	inits.InitRedis()
	api.LoadRouters(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
