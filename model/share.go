package model

import (
	"User/common"

	_ "github.com/go-sql-driver/mysql" //导入mysql
)

// var (
// 	ShareData []Stick_data
// )

func Set_stick(title, data string) bool {

	common.DB.AutoMigrate(&Stick_data{})
	// //增加
	common.DB.Create(&Stick_data{Title: title, Data: data})

	return true
}

func Get_stick() []Stick_data {

	var s []Stick_data
	var shareData []Stick_data

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
