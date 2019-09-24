package controllers

import (
	"emlog-go-gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BlogList(c *gin.Context) {
}

func Blog(c *gin.Context) {
	gid, _ := strconv.ParseInt(c.Param("gid"), 10, 64)
	fmt.Println(gid)
	blog := models.GetBlogById(gid)
	c.JSON(http.StatusOK, blog)
}