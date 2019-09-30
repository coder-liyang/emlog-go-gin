package controllers

import (
	"emlog-go-gin/common"
	"emlog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//分类列表
func SortRows(c *gin.Context) {
	sorts := models.GetSortList(0)
	c.JSON(http.StatusOK, common.GetResponse(0, sorts, "OK"))
}
