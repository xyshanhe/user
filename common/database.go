package common

import "github.com/jinzhu/gorm"

var (
	sql  = "root:password@tcp(localhost:3306)/project?charset=utf8&parseTime=True&loc=Local"
	DB   *gorm.DB
)


// Getsql 链接数据库
func Getslq() (err error) {
	DB, err = gorm.Open("mysql", sql)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
