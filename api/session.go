package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"yizu/modules"
	"yizu/service"
	yizuutil "yizu/util"
)

type SessionManager struct {
}

// AuthCode 校验验证码
func (*SessionManager) AuthCode(c *gin.Context) {
	phoneNum := c.PostForm("phone_num")
	code := c.PostForm("code")
	if phoneNum != "" && code != "" {
		redis := yizuutil.GetRedis()
		defer redis.Close()
		x, e := redis.Get(redis.Context(), phoneNum).Result()
		if e != nil || x != code {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "验证码错误",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "验证成功",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
	}
}

// GetAuthCode 获取短信验证码
// 手机号可能不合法
func (*SessionManager) GetAuthCode(c *gin.Context) {
	phoneNum := c.PostForm("phone_num")
	if phoneNum != "" {
		code := yizuutil.GenerateCode(6)
		ok := yizuutil.SendAuthCode(phoneNum, code)
		if ok {
			redis := yizuutil.GetRedis()
			defer redis.Close()
			ctx := redis.Context()
			redis.Set(ctx, phoneNum, code, 6*time.Minute)
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "验证码发送成功",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1,
				"msg":  "验证码发送失败, 请检查手机号是否正确或稍后再试",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
	}
}

// Login 用户登陆
// 1.采用手机加验证码方式登陆
// 2.采用用户名加密码方式登陆
// TODO 登陆成功删除Redis中的手机号缓存
func (*SessionManager) Login(c *gin.Context) {
	phoneNum := c.PostForm("phone_num")
	code := c.PostForm("code")
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 采用手机号登陆
	if phoneNum != "" && code != "" && username == "" && password == "" {
		redis := yizuutil.GetRedis()
		defer redis.Close()
		ctx := redis.Context()
		val, err := redis.Get(ctx, phoneNum).Result()
		if err != nil {
			c.JSON(http.StatusBadRequest, modules.ArgErr())
		} else {
			if val == code { // 登陆成功
				db, err := yizuutil.GetDB()
				if err != nil {
					c.JSON(http.StatusInternalServerError, modules.SysErr())
					return
				}
				userInfo := &modules.User{}
				num := db.Where(modules.User{Phone: phoneNum}).First(userInfo).RowsAffected
				if num == 0 {
					c.JSON(http.StatusBadRequest, modules.LoginFail())
				} else {
					key := yizuutil.GenerateCode(20)
					ok := service.LoginSuccess(key, userInfo)
					if ok {
						c.SetCookie("session.id", key, 2592000, "/", "", false, false)
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
			key := yizuutil.GenerateCode(20)
			ok := service.LoginSuccess(key, userInfo)
			if ok {
				c.SetCookie("session.id", key, 2592000, "/", "", false, false)
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
	key, err := c.Cookie("session.id")
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.Failure())
	} else {
		rdb := yizuutil.GetRedis()
		ctx := rdb.Context()
		rdb.Del(ctx, key)
		cookie, _ := c.Cookie("session.id")
		c.SetCookie("session.id", cookie, -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, modules.Success())
	}
}

// Register 用户注册
// 先获取验证码
func (*SessionManager) Register(c *gin.Context) {
	var user modules.RegistInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
		return
	}
	// 这里可以为了提升网络传输速度单独搞一个验证短信验证码的接口
	redis := yizuutil.GetRedis()
	code, _ := redis.Get(redis.Context(), user.Phone).Result()
	if code != user.AuthCode {
		c.JSON(http.StatusUnauthorized, modules.ResInfo{
			Code: 1,
			Msg:  "验证码错误",
		})
		return
	}
	// TODO 头像上传有问题
	//file, err := c.FormFile("picture")
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, modules.ArgErr())
	//	return
	//}
	//err = c.SaveUploadedFile(file, conf.ServerConfig().AvatarUrl)
	//if err != nil {
	//	log.Errorf("用户头像保存失败: %v", err)
	//}
	//user.Picture = conf.ServerConfig().AvatarUrl + file.Filename
	ok := service.RegisterUser(&user)
	if ok {
		c.JSON(http.StatusOK, modules.Success())
	} else {
		c.JSON(http.StatusBadRequest, modules.Failure())
	}

}

// Logoff 注销-将用户里的存活状态改成注销
// 不删除与用户相关联的所有数据
func (*SessionManager) Logoff(c *gin.Context) {
	key, err := c.Cookie("session.id")
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.Failure())
	} else {
		rdb := yizuutil.GetRedis()
		ctx := rdb.Context()
		k, _ := rdb.Get(ctx, key).Bytes()
		if k == nil {
			c.JSON(http.StatusBadRequest, modules.ArgErr())
			return
		}
		cacheInfo := modules.CacheInfo{}
		json.Unmarshal(k, &cacheInfo)
		db, err := yizuutil.GetDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, modules.Failure())
			return
		}
		err = db.Where(modules.User{Id: cacheInfo.UserId}).Update("status", 2).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, modules.Failure())
			return
		}
		rdb.Del(ctx, key)
		c.JSON(http.StatusOK, modules.Success())
	}
}
