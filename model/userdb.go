package model

/*
-----------------------------------
-----------------------------------
用户数据库信息（注册，更改，查询登录）
-----------------------------------
-----------------------------------
*/

import (
	"User/common"
	"User/util"

	_ "github.com/go-sql-driver/mysql" //导入mysql
)

var (
	UserInfo []User
)


// Set 增
func Set(n string, p string, m string, an string) bool {

	//创建表自动迁移
	common.DB.AutoMigrate(&User{})

	hash, err := util.HashEncode(p)
	util.Check(err)
	// db.Commit()

	// //增加
	common.DB.Create(&User{User_name: n, User_password: hash, User_mail: m, User_account_name: an})
	return true
}

//Get 查
func Get(n string, p string) bool {
	var u []User

	common.DB.Where("user_name = ?", n).First(&u)

	for _, j := range u {
		if util.ComparePwd(j.User_password, p) == true {
			return true
		}
	}
	return false
}

// Update 改密码
func Update(p string, m string) bool {

	var u []User

	hash, err := util.HashEncode(p)
	util.Check(err)

	common.DB.Model(&u).Where("user_mail = ?", m).Update("user_password", hash)

	return true
}

// UpdateAccountName 更改用户名称
func UpdateAccountName(user, name string) bool {
	var u []User

	common.DB.Model(&u).Where("user_name = ?", user).Update("user_account_name", name)

	return true
}

// MaliLogin 邮箱登录
func MaliLogin(m string) bool {

	var u []User

	common.DB.Where("user_mail = ? ", m).Find(&u)

	for _, j := range u {
		if j.User_mail == m {
			return true
		}
	}
	return false
}
