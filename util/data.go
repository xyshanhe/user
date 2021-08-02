package util

import "github.com/jinzhu/gorm"

// 用户数据
type UserInfo struct {
	// 定义结构体去前端绑定参数 form就写form json就写json
	gorm.Model
	Id       int    `form:"id" json:"id"`
	UserName string `form:"username" json:"username"`
	PassWord string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
	Uname    string `form:"uname" json:"uname"`
	YZM      string `form:"yzm" json:"yzm"`
	Code     string `form:"code" json:"code"`
	// var user_info Userinfo
	// err := c.ShouldBind(&结构体) 获取前端内容
}

//主页数据
type Code_data struct {
	Appname string `form:"appname" json:"appname"`
	Explain string `form:"explain" json:"explain"`
	Addr    string `form:"addr" json:"addr"`
	Imgs    string `form:"imgs" json:"imgs"`
}

//发帖数据
type Stick_data struct {
	Title string `form:"title" json:"title"`
	Data  string `form:"data" json:"data"`
}
