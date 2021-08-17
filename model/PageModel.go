package model

type Data struct {
	//gorm.Model
	Id       int
	Appname  string
	Explain  string
	Addr     string
	Imgs     string
	Category string
}
