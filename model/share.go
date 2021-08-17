package model

import (
	"User/common"
	_ "github.com/go-sql-driver/mysql" //导入mysql
)

var (
	ShareData []Stick_data
)

func Set_stick(title, data string) bool {

	common.DB.AutoMigrate(&Stick_data{})
	// //增加
	common.DB.Create(&Stick_data{Title: title, Data: data})

	return true
}

func Get_stick() {

	var s []Stick_data

	common.DB.Find(&s)

	ShareData = append(ShareData, s...)
}
