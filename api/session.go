package api

import "github.com/gin-gonic/gin"

type SessionManager struct {

}

// Login 用户登陆
func (*SessionManager) Login(c *gin.Context) {

}

// Logout 用户退出
func (*SessionManager) Logout(c *gin.Context) {

}

// Register 用户注册
func (*SessionManager) Register(c *gin.Context) {

}

// Logoff 注销-将用户里的存活状态改成注销
// 不删除与用户相关联的所有数据
func (*SessionManager) Logoff(c *gin.Context) {

}