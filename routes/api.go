package routes

import (
	"emlog-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine) {
	api := r.Group("api/")
	//文章列表
	api.GET("blog/row/:gid", controllers.BlogRow)
	//一篇文章
	api.GET("blog/rows", controllers.BlogRows)
	//文章评论
	api.GET("comment/rows/:gid", controllers.CommentRows)
}