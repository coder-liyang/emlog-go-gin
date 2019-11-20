package controllers

import (
	"emlog-go-gin/common"
	"emlog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//标签列表
func TagRows(c *gin.Context) {
	tags := models.GetTagList()
	c.JSON(http.StatusOK, common.GetResponse(0, tags, "OK"))
}