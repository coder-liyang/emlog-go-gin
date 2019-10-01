package controllers

import (
	"emlog-go-gin/common"
	"emlog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LinkRows(c *gin.Context) {
	links := models.GetLinkList()
	c.JSON(http.StatusOK, common.GetResponse(0, links, ""))
	return
}
