package controllers

import (
	"emlog-go-gin/common"
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
	c.JSON(http.StatusOK, common.GetResponse(0, blogSlice, "OK"))
}

//一篇文章
func BlogRow(c *gin.Context) {
	gid, _ := strconv.ParseInt(c.Param("gid"), 10, 64)
	blog, err := models.GetBlogById(gid)
	if err != nil {
		c.JSON(http.StatusOK, common.GetResponse(0, nil, "OK"))
		return
	}

	blog.DateFmt = time.Unix(blog.Date, 0).Format("2006-01-02 15:04:05")

	response := struct {
		*models.Blog
		Hide    string `json:"hide,omitempty"`
		Checked string `json:"checked,omitempty"`
	}{
		Blog: &blog,
	}
	c.JSON(http.StatusOK, common.GetResponse(0, response, "OK"))
	return
}
