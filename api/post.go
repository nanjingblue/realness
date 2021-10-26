package api

import (
	"github.com/gin-gonic/gin"
	"realness/service"
)

func PostPush(c *gin.Context)  {
	var service service.PostPushService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Push(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
// ShowPost 根据 id 获取串
func ShowPost(c *gin.Context)  {
	var service service.ShowPostService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListPosts 获取所有的串
func ListPosts(c *gin.Context)  {
	var service service.ListPostService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// LikePost 喜欢某 id 的串
func LikePost(c *gin.Context)  {
	var service service.LikePostService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Like(c.Param("id"), c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}