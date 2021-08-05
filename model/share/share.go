package share

import (
	"github.com/jinzhu/gorm"
)

type Stick_data struct {
	gorm.Model
	Title string
	Data  string
}

var (
	Info []Stick_data
	sql  = "root:password@tcp(localhost:3306)/project?charset=utf8&parseTime=True&loc=Local"
)

func Set_stick(title, data string) bool {

	//用户名:密码@tcp(ip:port)/数据库名称?charset=utf8&paresTime=True&loc=Local
	db, err := gorm.Open("mysql", sql)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Stick_data{})
	// //增加
	db.Create(&Stick_data{Title: title, Data: data})

	return true
}

func Get_stick() {

	var s []Stick_data

	//用户名:密码@tcp(ip:port)/数据库名称?charset=utf8&paresTime=True&loc=Local
	db, err := gorm.Open("mysql", sql)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Find(&s)

	for _, j := range s {
		Info = append(Info, j)
	}
}
