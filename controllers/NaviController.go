package controllers

import (
	"emlog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func NaviRows(c *gin.Context) {
	comments := models.GetNaviList(0)
	c.JSON(http.StatusOK, comments)
}