package controllers

import (
	"emlog-go-gin/common"
	"emlog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//导航
func NaviRows(c *gin.Context) {
	navis := models.GetNaviList(0)
	c.JSON(http.StatusOK, common.GetResponse(0, navis, "OK"))
}
