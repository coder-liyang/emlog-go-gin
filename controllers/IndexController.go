package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {

}

func Index(c *gin.Context) {
	//c.String(http.StatusOK, "hello index")
	c.HTML(http.StatusOK, "index/index.html", gin.H{})
}