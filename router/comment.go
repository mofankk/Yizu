package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yizu/api"
)

func commentRouter(r *gin.Engine) {

	c := api.CommentManager{}

	// 评论模块
	g := r.Group("/comment")
	g.Handle(http.MethodPost, "/house", c.CommentForHouse) // 用户对房子的评论
	g.Handle(http.MethodGet, "/house", c.ListForHouse)   // 查看房子下的评论信息
	g.Handle(http.MethodPost, "/user", c.CommentForUser)   // 用户间的评论
	g.Handle(http.MethodGet, "/user", c.ListForUser)    // 查看用户下的评论信息
}
