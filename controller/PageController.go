package controller

import (
	"User/reptile"
	"User/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Page 主页逻辑
func To_Page(c *gin.Context) {
	c.HTML(200, "page/index.html", nil)
}

func Do_Page(c *gin.Context) {
	var code util.Code_data

	err := c.ShouldBind(&code)
	util.Check(err)

	c.JSON(200, gin.H{
		"data":    content(),
		"status":  200,
		"message": "数据获取成功",
	})
}

func content() []map[string]string {
	var (
		slice []map[string]string
		arr1  map[string]string
	)

	reptile.Data_get()
	for _, i := range reptile.Info {
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