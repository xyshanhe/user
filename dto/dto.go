package dto

import "User/model/usermodel"

type UserDto struct {
	Name        string `json:"name,omitempty"`
	AccountName string `json:"AccountName,omtiempty"`
}

func ToUserDto(user usermodel.User) UserDto {
	return UserDto{
		Name:        user.User_name,
		AccountName: user.User_account_name,
	}
}