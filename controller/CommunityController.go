package controller

import (
	"User/model"
	"User/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


//community 社区主页
func To_community(c *gin.Context) {
	c.HTML(200, "community/index.html", nil)
}

func Do_community(c *gin.Context) {
	var code util.Stick_data

	err := c.ShouldBind(&code)
	util.Check(err)

	c.JSON(200, gin.H{
		"data":    com_content(),
		"status":  200,
		"message": "数据获取成功",
	})
}

func com_content() []map[string]interface{} {
	var (
		slice []map[string]interface{}
		arr1  map[string]interface{}
	)

	model.Get_stick()
	for _, i := range model.ShareData {
		arr1 = make(map[string]interface{})
		arr1["title"] = i.Title
		arr1["data"] = i.Data
		arr1["id"] = fmt.Sprint(i.ID)
		arr1["createtime"] = (i.CreatedAt)

		slice = append(slice, arr1)
	}
	return slice

}

// share 分享内容
func To_share(c *gin.Context) {
	c.HTML(200, "share/index.html", nil)
}

func Do_share(c *gin.Context) {
	var code util.Stick_data

	err := c.ShouldBind(&code)
	util.Check(err)

	if len(code.Title) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "标题不可为空！！！"})
		return
	}

	map_data := map[string]interface{}{
		"code": 200,
		"msg":  "分享成功"}

	if model.Set_stick(code.Title, code.Data) == true {
		c.JSON(http.StatusOK, map_data)
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "分享失败"})
	}
}
