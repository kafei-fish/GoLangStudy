package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("usersersion", "userid-1")
		context.Set("usersersion", "userid-1")
		context.Set("usersersion", "userid-1")
		context.Next()
	}
}
func main() {
	// 创建服务
	ginServer := gin.Default()

	// 加载静态页面
	ginServer.LoadHTMLGlob("admin/static/templates/*")
	//相应一个页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "go 传递的数据"})
	})
	// 接受前端传参 http://127.0.0.1:7899/user/info?userId=123&username=dfa
	ginServer.GET("/user/info", myHandler(), func(context *gin.Context) {
		// 拿到自定义中间件值
		usersession := context.MustGet("usersersion").(string)
		log.Println("============", usersession)
		userId := context.Query("userId")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userId,
			"username": username,
		})
	})
	// resultful 传参 http://127.0.0.1:7899/user/info/11/aa
	ginServer.GET("/user/info/:userId/:username", func(context *gin.Context) {
		// 拿到地址栏传递的参数
		userId := context.Param("userId")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userId":   userId,
			"userName": username,
		})
	})

	// json 前端传递给后端
	ginServer.POST("/json", func(context *gin.Context) {
		data, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})

	// 接受表单从传参
	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		log.Println("用户名：", username)
		log.Println("密码：", password)
	})

	// 路由组
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.POST("/login")
		userGroup.POST("/logout")
	}

	orderGroup := ginServer.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.DELETE("delete")
	}

	ginServer.Run(":7899")
}
