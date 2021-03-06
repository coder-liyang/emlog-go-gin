package routes

import (
	"emlog-go-gin/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterApi(r *gin.Engine) {
	r.GET("ws", controllers.WsHandler)
	//r.Use(Cors())
	api := r.Group("api").Use(Cors())
	//一篇文章
	api.GET("blog/row/:gid", controllers.BlogRow)
	//文章列表
	api.GET("blog/rows", controllers.BlogRows)
	//文章评论列表
	api.GET("comment/rows/:gid", controllers.CommentRows)
	//分类列表
	api.GET("sort/rows", controllers.SortRows)
	//导航列表
	api.GET("navi/rows", controllers.NaviRows)
	//友情链接列表
	api.GET("link/rows", controllers.LinkRows)
	//发表评论
	api.POST("comment/create", controllers.CommentCreate)
	//标签列表
	api.GET("tag/rows", controllers.TagRows)
	//发送邮件
	api.POST("mail/send", controllers.Send)
	//文章归档列表
	api.GET("blog/record", controllers.BlogRecord)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}