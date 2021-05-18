package handler

import (
	"github.com/gin-gonic/gin"
)

// 评价管理

type CommentManager struct {

}

// List 获取评论列表
// 用户浏览到房子，查看完详情后，可以向房东请求展示评论详情
//
func (*CommentManager) List(c *gin.Context) {

}

// Update 更新评论
func (*CommentManager) Update(c *gin.Context) {

}

// Delete 删除评论-只有管理员才能做
func (*CommentManager) Delete(c *gin.Context) {

}

// Add 新增评论
func (*CommentManager) Add(c *gin.Context) {

}
