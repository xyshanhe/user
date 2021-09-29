package model

import "github.com/jinzhu/gorm"

type StickData struct {
	gorm.Model
	Title string
	Data  string
}