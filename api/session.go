package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"yizu/modules"
	"yizu/service"
	yizuutil "yizu/util"
)

type SessionManager struct {

}

// GetAuthCode 获取短信验证码
// 手机号可能不合法
func (*SessionManager) GetAuthCode(c *gin.Context) {
	phoneNum := c.PostForm("phone_num")
	if phoneNum != "" {
		code := yizuutil.GenerateCode(6)
		ok := yizuutil.SendAuthCode(phoneNum, code)
		if ok {
			redisStore := yizuutil.GetRedisStore()
			conn := redisStore.Pool.Get()
			defer conn.Close()
			conn.Do("SET", phoneNum, code)
			conn.Do("EXPIRE", phoneNum, 6 * time.Minute)
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg": "验证码发送成功",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1,
				"msg": "验证码发送失败, 请检查手机号是否正确或稍后再试",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
	}
}

// Login 用户登陆
// 1.采用手机加验证码方式登陆
// 2.采用用户名加密码方式登陆
func (*SessionManager) Login(c *gin.Context) {
	phoneNum := c.PostForm("phone_num")
	code := c.PostForm("auth_code")
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 采用手机号登陆
	if phoneNum != "" && code != "" && username == "" && password == ""{
		redisStore := yizuutil.GetRedisStore()
		conn := redisStore.Pool.Get()
		defer conn.Close()
		cc, err := conn.Do("GET", phoneNum)
		if err != nil {
			c.JSON(http.StatusBadRequest, modules.ArgErr())
		} else {
			if cc.(string) == code { // 登陆成功
				db, err := yizuutil.GetDB()
				if err != nil {
					c.JSON(http.StatusInternalServerError, modules.SysErr())
					return
				}
				userInfo := &modules.User{}
				num := db.Where(modules.User{Phone: phoneNum}).First(userInfo).RowsAffected
				if num == 0 {

				} else {
					ok := service.LoginSuccess(userInfo)
					if ok {
						c.JSON(http.StatusOK, modules.LoginSuccess())
					} else {
						c.JSON(http.StatusOK, modules.LoginFail())
					}
				}
			} else {
				c.JSON(http.StatusBadRequest, modules.LoginFail())
			}
		}
	} else if username != "" && password != "" { // 采用用户名和密码方式登陆
		if username == "" || password == "" {
			c.JSON(http.StatusBadRequest, modules.ArgErr())
			return
		}
		db, err := yizuutil.GetDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, modules.SysErr())
			return
		}
		userInfo := &modules.User{}
		num := db.Where(modules.User{Username: username, Password: password}).First(userInfo).RowsAffected
		if num == 0 {
			c.JSON(http.StatusInternalServerError, modules.ResInfo{
				1,
				"用户名或密码不正确",
			})
			return
		} else { //登陆成功
			ok := service.LoginSuccess(userInfo)
			if ok {
				c.JSON(http.StatusOK, modules.LoginSuccess())
			} else {
				c.JSON(http.StatusOK, modules.LoginFail())
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
	}
}

// Logout 用户退出
func (*SessionManager) Logout(c *gin.Context) {

}

// Register 用户注册
// 先获取验证码
func (*SessionManager) Register(c *gin.Context) {

}

// Logoff 注销-将用户里的存活状态改成注销
// 不删除与用户相关联的所有数据
func (*SessionManager) Logoff(c *gin.Context) {

}