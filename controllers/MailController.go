package controllers

import (
	"emlog-go-gin/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Send(c *gin.Context)  {
	go SendMailDemo()
	c.JSON(http.StatusOK, common.GetResponse(0, nil, "OK"))
}

func SendMailDemo() {
	time.Sleep(time.Duration(5)*time.Second)
	fmt.Println("send mail demo func")


}
