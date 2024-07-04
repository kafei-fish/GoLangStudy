package api

import (
	"GoLangStudy/model/request"
	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(ctx gin.Context) {
	var registerParam request.RegisterParam
	// 转换josn
	_ = ctx.ShouldBindJSON(registerParam)
	// 注册
	//register, err := service.Register(registerParam)
	//// 如果错误信息不空
	//if err != nil {
	//	ctx.JSON(200, gin.H{
	//		"code": "500",
	//		"msg":  "注册失败",
	//	})
	//	return
	//}
}
