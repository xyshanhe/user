package model

import "github.com/jinzhu/gorm"

type Stick_data struct {
	gorm.Model
	Title string
	Data  string
}