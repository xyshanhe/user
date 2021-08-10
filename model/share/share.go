package share

import (
	"User/model"

	_ "github.com/go-sql-driver/mysql" //导入mysql
	"github.com/jinzhu/gorm"
)

type Stick_data struct {
	gorm.Model
	Title string
	Data  string
}

var (
	Info []Stick_data
)

func Set_stick(title, data string) bool {

	model.DB.AutoMigrate(&Stick_data{})
	// //增加
	model.DB.Create(&Stick_data{Title: title, Data: data})

	return true
}

func Get_stick() {

	var s []Stick_data

	model.DB.Find(&s)

	Info = append(Info, s...)
}
