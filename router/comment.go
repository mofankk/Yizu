package router

import (
	"github.com/gin-gonic/gin"
	"yizu/api"
)

func commentRouter(r *gin.Engine) {

	c := api.CommentManager{}

	// 评论模块
	g := r.Group("/comment")
	g.Handle("POST", "/add/house", c.CommentForHouse) // 用户对房子的评论
	g.Handle("POST", "/add/user", c.CommentForUser)   // 用户间的评论

}
