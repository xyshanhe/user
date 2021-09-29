package util

import "github.com/jinzhu/gorm"

// UserInfo 用户数据
type UserInfo struct {
	gorm.Model
	Id       int    `form:"id" json:"id"`
	UserName string `form:"username" json:"username"`
	PassWord string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
	Uname    string `form:"uname" json:"uname"`
	YZM      string `form:"yzm" json:"yzm"`
	Code     string `form:"code" json:"code"`
}

// CodeData app内容
type CodeData struct {
	Appname string `form:"appname" json:"appname"`
	Explain string `form:"explain" json:"explain"`
	Addr    string `form:"addr" json:"addr"`
	Imgs    string `form:"imgs" json:"imgs"`
}

// StickData 发帖
type StickData struct {
	Title string `form:"title" json:"title"`
	Data  string `form:"data" json:"data"`
}
