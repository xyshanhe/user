package usermodel

//模型大写 首字母大写
type User struct {
	//自带四个字段：ID/更新时间/创建时间/删除时间/
	// gorm.Model
	Id                int    `gorm:"primary_key"`
	User_name         string `gorm:"type:varchar(64)"`
	User_password     string `gorm:"type:varchar(32)"`
	User_mail         string `gorm:"type:varchar(128)"`
	User_account_name string `gorm:"type:varchar(32)"`
	// user_update_time
	// user_create_time
}
