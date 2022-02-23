package controller

import (
	"User/model"
	"User/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ToCommunity 社区主页
func ToCommunity(c *gin.Context) {
	c.HTML(200, "community/index.html", nil)
}

func DoCommunity(c *gin.Context) {
	var code util.StickData

	err := c.ShouldBind(&code)
	util.Check(err)

	c.JSON(200, gin.H{
		"data":    comContent(),
		"status":  200,
		"message": "数据获取成功",
	})
}

func comContent() []map[string]interface{} {
	var (
		slice []map[string]interface{}
		arr1  map[string]interface{}
	)

	var shareData = model.GetStick()

	for _, i := range shareData {
		arr1 = make(map[string]interface{})
		arr1["title"] = i.Title
		arr1["data"] = i.Data
		arr1["id"] = fmt.Sprint(i.ID)
		arr1["createtime"] = (i.CreatedAt)

		slice = append(slice, arr1)
	}
	return slice

}

// ToShare 分享内容
func ToShare(c *gin.Context) {
	c.HTML(200, "share/index.html", nil)
}

func DoShare(c *gin.Context) {
	var code util.StickData

	err := c.ShouldBind(&code)
	util.Check(err)

	if len(code.Title) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "标题不可为空！！！"})
		return
	}

	if len(code.Data) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请输入了有效链接"})
		return
	}

	MapData := map[string]interface{}{
		"code": 200,
		"msg":  "分享成功"}

	if model.SetStick(code.Title, code.Data) == true {
		c.JSON(http.StatusOK, MapData)
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "分享失败"})
	}
}
