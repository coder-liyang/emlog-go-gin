package routes

import (
	"emlog-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine) {
	api := r.Group("api/")
	api.GET("blog/row/:gid", controllers.BlogRow)
	api.GET("blog/rows", controllers.BlogRows)
}