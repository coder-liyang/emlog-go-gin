package controllers

import (
	"emlog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//文章列表
func BlogRows(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	blogList := models.GetBlogList(page)
	blogSlice := make([]interface{}, 0, len(blogList))
	for _, item := range blogList {
		item.DateFmt = time.Unix(item.Date, 0).Format("2006-01-02 15:04:05")
		blogSlice = append(blogSlice, item)
	}
	c.JSON(http.StatusOK, blogSlice)
}

//一篇文章
func BlogRow(c *gin.Context) {
	gid, _ := strconv.ParseInt(c.Param("gid"), 10, 64)
	blog := models.GetBlogById(gid)
	c.JSON(http.StatusOK, blog)
}
