package controller

import (
	"GoLangStudy/domain"
	"GoLangStudy/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const USER_NAME = "admin"
const PASS_WORD = "123456"

/*
登录方法
*/
func HandleUserLogin(context *gin.Context) {
	name := context.PostForm("username")
	password := context.PostForm("password")
	po := domain.LoginPo{
		UserName: name,
		Password: password,
	}
	b := service.UserLogin(po)
	if b {
		context.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "登录成功",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": "500",
			"msg":  "登录失败",
		})
	}
}
