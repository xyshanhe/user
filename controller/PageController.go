package controller

import (
	"User/reptile"
	"User/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ToPage1(c *gin.Context) {
	c.HTML(200, "page1/page1.html", nil)
}

func DoPage1(c *gin.Context) {
	var code util.CodeData

	err := c.ShouldBind(&code)
	util.Check(err)

	c.JSON(200, gin.H{
		"data":    content(),
		"status":  200,
		"message": "数据获取成功",
	})
}

// ToPage 主页逻辑
func ToPage(c *gin.Context) {
	c.HTML(200, "page/index.html", nil)
}

// DoPage 获取数据主页数据
func DoPage(c *gin.Context) {
	var code util.CodeData

	err := c.ShouldBind(&code)
	util.Check(err)

	c.JSON(200, gin.H{
		"data":    content(),
		"status":  200,
		"message": "数据获取成功",
	})
}

// content 主页数据处理
func content() []map[string]string {
	var (
		slice []map[string]string
		arr1  map[string]string
	)

	var data = reptile.DataGet()

	for _, i := range data {
		arr1 = make(map[string]string)
		arr1["appname"] = i.Appname
		arr1["Exp"] = i.Explain
		arr1["Addr"] = i.Addr
		arr1["logo"] = i.Imgs
		arr1["Id"] = fmt.Sprint(i.Id)
		arr1["Category"] = i.Category

		slice = append(slice, arr1)
	}

	return slice
}
