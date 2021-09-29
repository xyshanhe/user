package main

import (
	"User/controller"
	"User/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {



	r.LoadHTMLGlob("template/**/*")
	r.Static("/static", "./static")

	//r.Use(middleware.AuthMiddleware()) //用户token中间件
	//r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware()) //跨域跨域
	catgoryRoutes := r.Group("/user")

	//catgoryRoutes := r.Group("/user",middleware.AuthMiddleware())
	//{
	//
	//}
	catgoryRoutes.GET("/to/register", controller.ToRegister)
	catgoryRoutes.POST("/do/register", controller.DoRegister)
	//登录
	catgoryRoutes.GET("/to/login", controller.ToLogin)
	catgoryRoutes.POST("/do/login", controller.DoLogin)
	//修改密码
	catgoryRoutes.GET("/to/update", controller.ToUpdate)
	catgoryRoutes.POST("/do/update", controller.DoUpdate)

	catgoryRoutes.POST("/emailyzm", controller.EmailYzm)
	catgoryRoutes.POST("/email", controller.EmailLogin)
	//修改用户名称
	catgoryRoutes.PUT("/UpdateAccountName", controller.UpdateAccountName)

	catgoryRoutes.GET("/info", middleware.AuthMiddleware(), controller.Info)



	//主页
	r.GET("/to/page", controller.ToPage)
	r.GET("/do/page", controller.DoPage)

	r.GET("/to/page1", controller.ToPage1)
	r.GET("/do/page1", controller.DoPage1)

	//社区
	r.GET("to/community", controller.ToCommunity)
	r.GET("do/community", controller.DoCommunity)
	r.GET("/community/to/share", controller.ToShare)
	r.POST("/community/do/share", controller.DoShare)

	return r
}
