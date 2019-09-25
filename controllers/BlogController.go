package controllers

import (
	"emlog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BlogRows(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	blogList := models.GetBlogList(page)
	c.JSON(http.StatusOK, blogList)
}

func BlogRow(c *gin.Context) {
	gid, _ := strconv.ParseInt(c.Param("gid"), 10, 64)
	blog := models.GetBlogById(gid)
	c.JSON(http.StatusOK, blog)
}