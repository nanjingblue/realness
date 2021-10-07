package server

import (
	"os"
	"realness/api"
	"realness/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)
		user := v1.Group("/user")
		{
			user.POST("/register", api.UserRegister)
			user.POST("/login", api.UserLogin)
			// 需要登录保护的
			auth := user.Group("")
			auth.Use(middleware.AuthRequired())
			{
				// User Routing
				auth.GET("/me", api.UserMe)
				auth.DELETE("/logout", api.UserLogout)
			}
		}
		post := v1.Group("")
		{
			post.GET("/posts", api.ListPosts)
			post.GET("/post/:id", api.ShowPost)
			// 需要登录保护
			auth := post.Group("/post")
			auth.Use(middleware.AuthRequired())
			{
				auth.POST("", api.PostPush)
				auth.GET("/like/:id", api.LikePost)
			}
		}
	}
	return r
}
