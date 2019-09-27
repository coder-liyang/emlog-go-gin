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
	keyword := c.Query("keyword")
	sortid, err := strconv.ParseInt(c.Query("sortid"), 10, 64)
	if err != nil {
		sortid = 0
	}
	blogList := models.GetBlogList(page, keyword, sortid)
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
	blog.DateFmt = time.Unix(blog.Date, 0).Format("2006-01-02 15:04:05")

	//c.JSON(http.StatusOK, blog)
	c.JSON(http.StatusOK, struct {
		*models.Blog
		Hide    string `json:"hide,omitempty"`
		Checked string `json:"checked,omitempty"`
	}{
		Blog: &blog,
	})
}
