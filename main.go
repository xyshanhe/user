package main

import (
	"User/middleware"
	"User/model"

	"github.com/gin-gonic/gin"
)

func main() {

	err := model.Getslq()
	if err != nil {
		panic(err)
	}

	defer model.DB.Close()

	//创建路由
	r := gin.Default()
	r.LoadHTMLGlob("template/**/*")
	r.Static("/static", "./static")

	//创建
	r.GET("/user/to/register", To_Register)
	r.POST("/user/do/register", Do_Register)

	//登录
	r.GET("/user/to/login", To_Login)
	r.POST("/user/do/login", Do_Login)
	r.GET("/user/info", middleware.AuthMiddleware(), Info)

	//修改用户名称
	r.PUT("/user/UpdateAccountName", UpdateAccountName)

	//修改密码
	r.GET("/user/to/update", To_update)
	r.POST("/user/do/update", Do_update)

	r.POST("/user/emailyzm", Email_Yzm)
	r.POST("/user/email", Email_login)

	//主页
	r.GET("/to/page", To_Page)
	r.GET("/do/page", Do_Page)

	//社区
	r.GET("to/community", To_community)
	r.GET("do/community", Do_community)
	r.GET("/community/to/share", To_share)
	r.POST("/community/do/share", Do_share)
	//开启监听

	r.Run(":8080")
}
