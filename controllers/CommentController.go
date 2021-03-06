package controllers

import (
	"emlog-go-gin/common"
	"emlog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//评论列表
func CommentRows(c *gin.Context) {
	gid, _ := strconv.ParseInt(c.Param("gid"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	comments := models.GetCommentList(gid, 0, page, 10)
	c.JSON(http.StatusOK, common.GetResponse(0, comments, "OK"))
}

//发表评论
func CommentCreate(c *gin.Context) {
	gid, _ := strconv.ParseInt(c.Query("gid"), 10, 64)
	pid, _ := strconv.ParseInt(c.Query("pid"), 10, 64)
	if gid == 0 {
		c.JSON(http.StatusOK, common.GetResponse(10, nil, "gid为必填"))
		return
	}
	content, hasContent := c.GetPostForm("content")
	if hasContent == false {
		c.JSON(http.StatusOK, common.GetResponse(10, nil, "评论内容不可为空"))
		return
	}
	blog, err := models.Blog{}.GetBlogById(gid)
	if err != nil {
		c.JSON(http.StatusOK, common.GetResponse(10, nil, "文章不存在"))
		return
	}
	if blog.AllowRemark == "n" {
		c.JSON(http.StatusOK, common.GetResponse(11, nil, "不可评论"))
		return
	}
	if pid != 0 {
		pComment, err := models.GetComment(pid)
		if err != nil {
			c.JSON(http.StatusOK, common.GetResponse(10, nil, "父评论不存在"))
			return
		}
		if pComment.Gid != gid {
			c.JSON(http.StatusOK, common.GetResponse(10, nil, "父评论与文章不匹配"))
			return
		}
	}
	comment := models.Comment{
		Gid:           gid,
		Pid:           pid,
		Date:          time.Now().Unix(),
		Poster:        "",
		Comment:       content,
		Mail:          "",
		Url:           "",
		Ip:            "",
		Hide:          "n",
		ChildComments: nil,
	}
	err = models.CreateComment(&comment)
	if err != nil {
		c.JSON(http.StatusOK, common.GetResponse(10, nil, "评论发表失败"))
	}
	c.JSON(http.StatusOK, common.GetResponse(0, comment, "OK"))
}