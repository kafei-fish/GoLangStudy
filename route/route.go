package route

import (
	"GoLangStudy/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	ginServer := gin.Default()

	// 加载静态页面
	ginServer.LoadHTMLGlob("admin/static/templates/*")
	// 加载路由
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "go 传递的数据"})
	})
	// 将 POST 请求的处理逻辑抽取到 handlers 包中的 HandleUserLogin 函数
	ginServer.POST("user/userLogin", controller.HandleUserLogin)

	return ginServer
}
