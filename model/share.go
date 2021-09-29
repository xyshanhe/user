package model

import (
	"User/common"
	_ "github.com/go-sql-driver/mysql" //导入mysql
)

// SetStick 插入分享帖子
func SetStick(title, data string) bool {

	common.DB.AutoMigrate(&StickData{})
	// //增加
	common.DB.Create(&StickData{Title: title, Data: data})

	return true
}

// GetStick 获取分享帖子数据
func GetStick() []StickData {

	var s []StickData
	var shareData []StickData

	common.DB.Find(&s)

	//反转数据
	lengtg := len(s)
	for i := 0; i < lengtg/2; i++ {
		temp := s[lengtg-1-i]
		s[lengtg-1-i] = s[i]
		s[i] = temp
	}

	shareData = append(shareData, s...)

	return shareData
}
