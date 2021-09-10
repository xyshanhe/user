package main

import (
	"User/controller"
	"User/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	//r.Use(middleware.AuthMiddleware())

	r.LoadHTMLGlob("template/**/*")
	r.Static("/static", "./static")

	catgoryRoutes := r.Group("/user")

	catgoryRoutes.GET("/to/register", controller.ToRegister)
	catgoryRoutes.POST("/do/register", controller.DoRegister)

	//登录
	catgoryRoutes.GET("/to/login", controller.ToLogin)
	catgoryRoutes.POST("/do/login", controller.DoLogin)
	catgoryRoutes.GET("/info", middleware.AuthMiddleware(), controller.Info)

	//修改用户名称
	catgoryRoutes.PUT("/UpdateAccountName", controller.UpdateAccountName)

	//修改密码
	catgoryRoutes.GET("/to/update", controller.ToUpdate)
	catgoryRoutes.POST("/do/update", controller.DoUpdate)

	catgoryRoutes.POST("/emailyzm", controller.EmailYzm)
	catgoryRoutes.POST("/email", controller.EmailLogin)

	//主页
	r.GET("/to/page", controller.To_Page)
	r.GET("/do/page", controller.Do_Page)

	r.GET("/to/page1", controller.To_Page1)
	r.GET("/do/page1", controller.Do_Page1)

	//社区
	r.GET("to/community", controller.To_community)
	r.GET("do/community", controller.Do_community)
	r.GET("/community/to/share", controller.To_share)
	r.POST("/community/do/share", controller.Do_share)

	return r
}
