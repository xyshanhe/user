package model

/*
-----------------------------------
-----------------------------------
用户数据库信息（注册，更改，查询登录）
-----------------------------------
-----------------------------------
*/

import (
	"User/model/usermodel"
	"User/util"

	_ "github.com/go-sql-driver/mysql" //导入mysql
	"github.com/jinzhu/gorm"
)

var (
	Info []usermodel.User
	sql  = "root:password@tcp(localhost:3306)/project?charset=utf8&parseTime=True&loc=Local"
	DB   *gorm.DB
)

//链接数据库
func Getslq() (err error) {
	DB, err = gorm.Open("mysql", sql)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

// Set 增
func Set(n string, p string, m string, an string) bool {

	//创建表自动迁移
	DB.AutoMigrate(&usermodel.User{})

	hash, err := util.HashEncode(p)
	util.Check(err)
	// db.Commit()

	// //增加
	DB.Create(&usermodel.User{User_name: n, User_password: hash, User_mail: m, User_account_name: an})
	return true
}

//Get 查
func Get(n string, p string) bool {
	var u []usermodel.User

	DB.Where("user_name = ?", n).First(&u)

	for _, j := range u {
		if util.ComparePwd(j.User_password, p) == true {
			return true
		}
	}
	return false
}

// Update 改密码
func Update(p string, m string) bool {

	var u []usermodel.User

	hash, err := util.HashEncode(p)
	util.Check(err)

	DB.Model(&u).Where("user_mail = ?", m).Update("user_password", hash)

	return true
}

// UpdateAccountName 更改用户名称
func UpdateAccountName(user, name string) bool {
	var u []usermodel.User

	DB.Model(&u).Where("user_name = ?", user).Update("user_account_name", name)

	return true
}

// MaliLogin 邮箱登录
func MaliLogin(m string) bool {

	var u []usermodel.User

	DB.Where("user_mail = ? ", m).Find(&u)

	for _, j := range u {
		if j.User_mail == m {
			return true
		}
	}
	return false
}
