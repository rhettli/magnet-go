package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"magnet/glog"
	"magnet/controller"

)


//定义log文件
const MERCHANT_LOG="magnet"


//主svr router
func InitRouter() *gin.Engine {

	//初始化log
	glog.Init(MERCHANT_LOG)

	router := gin.Default()

	router.LoadHTMLGlob("tpl/*")
	router.Static("/static","static/")

	//初始化Group&log
	V0 := router.Group("/magnet")//api路由
	{

		//V0.GET("/hash", controller.SearchHash)

		V0.HEAD("/search/:key/:page", controller.SearchHashApi)

		//V0.GET("/detail/:hashId", controller.DetailPage)

	}

	router.GET("/", controller.HomePage)
	router.GET("/search/:key/:page", controller.SearchPage)//搜索页面展示
	router.GET("/detail/:id", controller.DetailPage)//详情页面展示
	router.GET("/about", controller.About)
	router.GET("/donate", controller.Donate)


	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	return router
}
