package main

import (
	"User/model"
	"User/model/share"
	"User/model/usermodel"
	"User/reptile"
	"User/sendmail"
	"User/util"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 注册
func To_Register(c *gin.Context) {
	c.HTML(200, "register/index.html", nil)
}

func Do_Register(c *gin.Context) {
	var user_info util.UserInfo
	err := c.ShouldBind(&user_info)
	util.Check(err)

	if len(user_info.UserName) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账号不能为空"})
		return
	}

	if len(user_info.PassWord) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "密码需要6位数"})
		return
	}

	if len(user_info.Email) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "邮箱不能为空"})
		return
	}

	if len(user_info.Uname) == 0 {
		user_info.Uname = util.RandomString(10)
	}

	map_data := map[string]interface{}{
		"code": 200,
		"msg":  "注册成功"}

	if user_info.YZM != sendmail.Code || user_info.YZM == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "验证码错误"})
		return
	}

	if model.Set(user_info.UserName, user_info.PassWord, user_info.Email, user_info.Uname) == false {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "用户名已经存在！！！"})
		return
	}

	model.Set(user_info.UserName, user_info.PassWord, user_info.Email, user_info.Uname)
	c.JSON(http.StatusOK, map_data)

}

// Login 登录
func To_Login(c *gin.Context) {
	c.HTML(200, "userlogin/login.html", nil)
}

func Do_Login(c *gin.Context) {
	var user_info util.UserInfo

	err := c.ShouldBind(&user_info)
	util.Check(err)
	// fmt.Println(user_info.UserName, user_info.PassWord)

	if len(user_info.UserName) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账号不能为空"})
		return
	}

	if len(user_info.PassWord) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码需要6位数"})
		return
	}

	token, err := util.ReleastToken(usermodel.User{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}

	map_data := map[string]interface{}{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功"}

	if model.Get(user_info.UserName, user_info.PassWord) == true {
		c.JSON(http.StatusOK, map_data)
	} else {
		c.String(http.StatusOK, "账号或密码错误")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "账号或密码错误"})
	}
}

//邮箱登录
func Email_login(c *gin.Context) {
	var user_info util.UserInfo
	err := c.ShouldBind(&user_info)
	util.Check(err)

	if len(user_info.Email) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请输入有效邮箱！！"})
		return
	}

	map_data := map[string]interface{}{
		"code": 200,
		"msg":  "验证码发送成功"}

	if model.MaliLogin(user_info.Email) == true {
		sendmail.DayMail(user_info.Email)
		c.JSON(http.StatusOK, map_data)
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "邮箱错误！！"})
	}

}

// UpdateAccountName 修改用户名称
func UpdateAccountName(c *gin.Context) {
	var user_info util.UserInfo
	err := c.ShouldBind(&user_info)
	util.Check(err)

	if len(user_info.UserName) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账号不能为空"})
		return
	}

	if len(user_info.Uname) == 0 {
		user_info.Uname = util.RandomString(10)
	}

	if model.UpdateAccountName(user_info.UserName, user_info.Uname) == true {
		c.JSON(200, gin.H{
			"msg": "修改成功",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "修改名称失败"})
	}
}

// To_update 修改密码
func To_update(c *gin.Context) {
	c.HTML(200, "updatapwd/index.html", nil)
}

func Do_update(c *gin.Context) {
	var user_info util.UserInfo
	err := c.ShouldBind(&user_info)
	util.Check(err)

	if len(user_info.PassWord) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "密码需要6位数"})
		return
	}

	if len(user_info.Email) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "邮箱不能为空"})
		return
	}

	if model.Update(user_info.PassWord, user_info.Email) == false {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "修改失败"})
		return
	}

	c.JSON(200, gin.H{
		"msg": "修改成功",
	})
}

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

	share.Get_stick()
	for _, i := range share.Info {
		arr1 = make(map[string]interface{})
		arr1["title"] = i.Title
		arr1["data"] = i.Data
		arr1["id"] = fmt.Sprint(i.Id)
		arr1["createtime"] = (i.Create_time)
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

	if share.Set_stick(code.Title, code.Data) == true {
		c.JSON(http.StatusOK, map_data)
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "分享失败"})
	}
}

//用户信息路由
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}
