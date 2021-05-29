package router

import (
	"github.com/gin-gonic/gin"
	"yizu/api"
)

func userRouter(r *gin.Engine) {

	u := api.UserManager{}

	// 用户模块
	g := r.Group("/user")
	g.Handle("GET", "/update", u.Update) // 修改用户信息，如果修改手机号需要先进行验证

	// 管理员操作
	g.Handle("POST", "/add", u.List)

}