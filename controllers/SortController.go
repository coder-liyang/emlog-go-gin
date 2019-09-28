package controllers

import (
	"emlog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func SortRows(c *gin.Context) {
	comments := models.GetSortList(0)
	c.JSON(http.StatusOK, comments)
}