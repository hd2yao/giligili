package server

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/hd2yao/giligili/api"
	"github.com/hd2yao/giligili/middleware"
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
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}

		// 视频投稿接口
		v1.POST("videos", api.CreateVideo)
		//// 视频详情接口
		//v1.GET("video/:id", api.ShowVideo)
		//// 视频列表接口
		//v1.GET("videos", api.ListVideo)
		//// 更新视频接口
		//v1.PUT("video/:id", api.UpdateVideo)
		//// 删除视频接口
		//v1.GET("video/:id", api.DeleteVideo)
	}
	return r
}