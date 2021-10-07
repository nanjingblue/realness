package service

import (
	"github.com/gin-gonic/gin"
	"realness/model"
	"realness/serializer"
)

// PostPushService 管理发串的服务
type PostPushService struct {
	Text string `form:"text" json:"text" binding:"required,min=1,max=1000"`
	Cover string `form:"cover" json:"cover" binding:"required,min=0,max=1000"`
}

// Push 当前用户发串函数
func (service *PostPushService) Push(c *gin.Context) serializer.Response {
	u, _ := c.Get("user")
	user, _ := u.(*model.User)
	post := model.Post {
		UserID: user.ID,
		Text: service.Text,
		Cover: service.Cover,
	}
	err := model.DB.Create(&post).Error
	if err != nil {
		return serializer.Response {
			Code: 501,
			Msg: "创建串失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildPost(post),
		Msg: "创建串成功",
	}
}

// ShowPostService 根据 id 获取串的服务
type ShowPostService struct {}

// Show  获取id串的函数
func (service *ShowPostService) Show(id string) serializer.Response {
	var post model.Post
	err := model.DB.First(&post, id).Error
	if err != nil {
		return serializer.Response{
			Code: 404,
			Msg: "视频不存在",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildPost(post),
		Msg: "获取串成功",
	}
}

// ListPostService 获取所有串的服务
type ListPostService struct {}

// List 获取串列表的函数
func (service *ListPostService) List() serializer.Response {
	var posts []model.Post
	err := model.DB.Find(&posts).Error
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg: "数据库连接错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildPosts(posts),
		Msg: "获取串列表成功",
	}
}

// LikePostService like 串的服务
type LikePostService struct {}

func (*LikePostService) Like(id string, c *gin.Context) serializer.Response {
	var post model.Post
	err := model.DB.First(&post, id).Error
	if err != nil {
		return serializer.Response{
			Code: 404,
			Msg: "视频不存在",
			Error: err.Error(),
		}
	}
	post.LikeNum++
	model.DB.Save(&post)

	u, _ := c.Get("user")
	user, _ := u.(*model.User)
	userLike := model.UserLike {
		UserID: user.ID,
		PostID: post.ID,
	}
	err = model.DB.Create(&userLike).Error
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg: "数据库连接失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: gin.H{"post": serializer.BuildPost(post), "userLike": serializer.BuildUserLike(userLike)},
		Msg: "喜欢串成功",
	}
}
