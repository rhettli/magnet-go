package main

import (

	"magnet/router"
	"github.com/gin-gonic/gin"
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("runing...","os:",runtime.GOOS)
	//启动svr，监听端口，对外提供服务
	gin.SetMode(gin.DebugMode)
	router := router.InitRouter()
	router.Run(":8088")
}
