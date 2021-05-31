package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yizu/api"
)

func sessionRouter(r *gin.Engine) {

	s := api.SessionManager{}

	r.Handle(http.MethodGet, "/code", s.GetAuthCode) // 获取验证码
	r.Handle(http.MethodPost, "/code", s.AuthCode)   // 校验验证码

	r.Handle(http.MethodPost, "/register", s.Register) // 用户注册
	r.Handle(http.MethodPost, "/login", s.Login)       // 用户登陆
	r.Handle(http.MethodDelete, "/logout", s.Logout)   // 退出登陆
	r.Handle(http.MethodDelete, "/logoff", s.Logoff)   // 用户注销
}