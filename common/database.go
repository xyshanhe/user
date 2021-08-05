// package common

// import (
// 	// _ "github.com/go-sql-driver/mysql" //导入mysql
// 	"github.com/jinzhu/gorm"
// )

// var (
// 	sql = "root:password@tcp(localhost:3306)/project?charset=utf8&parseTime=True&loc=Local"
// 	DB  *gorm.DB
// )

// func Getslq() (err error) {
// 	DB, err = gorm.Open("mysql", sql)
// 	if err != nil {
// 		return
// 	}
// 	return DB.DB().Ping()
// }
