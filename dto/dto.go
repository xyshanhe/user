package dto

import (
	"User/model"
)

type UserDto struct {
	Name        string `json:"name,omitempty"`
	AccountName string `json:"AccountName,omtiempty"`
}

// ToUserDto 返回格式
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:        user.User_name,
		AccountName: user.User_account_name,
	}
}
