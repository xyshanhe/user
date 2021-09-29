package controller

import (
	"User/common"
	"User/dto"
	"User/middleware"
	"User/model"
	"User/sendmail"
	"User/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ToRegister 注册
func ToRegister(c *gin.Context) {
	c.HTML(200, "register/index.html", nil)
}

func DoRegister(c *gin.Context) {
	var userInfo util.UserInfo
	err := c.ShouldBind(&userInfo)
	util.Check(err)

	if len(userInfo.UserName) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账号不能为空"})
		return
	}

	if len(userInfo.PassWord) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "密码需要6位数"})
		return
	}

	if len(userInfo.Email) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "邮箱不能为空"})
		return
	}

	//随机生成9位用户名称
	if len(userInfo.Uname) == 0 {
		userInfo.Uname = util.RandomString(9)
	}

	if userInfo.YZM != sendmail.Code || userInfo.YZM == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "验证码错误"})
		return
	}

	mapData := map[string]interface{}{
		"code": 200,
		"msg":  "注册成功"}

	if model.Set(userInfo.UserName, userInfo.PassWord, userInfo.Email, userInfo.Uname) == false {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "用户名已经存在！！！"})
		return
	}

	c.JSON(http.StatusOK, mapData)
}

// ToLogin 登录
func ToLogin(c *gin.Context) {
	c.HTML(200, "userlogin/login.html", nil)
}


func DoLogin(c *gin.Context) {
	var userInfo util.UserInfo
	var user model.User

	err := c.ShouldBind(&userInfo)
	util.Check(err)

	if len(userInfo.UserName) == 0 || len(userInfo.PassWord) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账号密码不能为空！！"})
		return
	}
	if len(userInfo.PassWord) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码需要6位数！！"})
		return
	}

	//查询用户是否存在
	common.DB.Where("user_name = ?", userInfo.UserName).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在！！！"})
		return
	}
	if model.Get(userInfo.UserName, userInfo.PassWord) == false {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误！！！"})
		return
	}

	//发放token
	token, err := middleware.ReleastToken(user)

	mapData := map[string]interface{}{"code": 200, "data": gin.H{"token": token, "username": user.User_account_name}, "msg": "登录成功"}

	//发送cookise
	//c.SetCookie("access_token", token, 60*3, "/", "127.0.0.1/user/", false, true)

	//返回结果
	c.JSON(http.StatusOK, mapData)
}

// EmailLogin 邮箱登录
func EmailLogin(c *gin.Context) {
	var userInfo util.UserInfo
	err := c.ShouldBind(&userInfo)
	util.Check(err)

	if len(userInfo.Email) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请输入有效邮箱！！"})
		return
	}

	MapData := map[string]interface{}{
		"code": 200,
		"msg":  "验证码发送成功"}

	if model.MaliLogin(userInfo.Email) == true {
		sendmail.DayMail(userInfo.Email)
		c.JSON(http.StatusOK, MapData)
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "邮箱错误！！"})
	}

}

// EmailYzm 发送验证码
func EmailYzm(c *gin.Context) {

	var userInfo util.UserInfo
	err := c.ShouldBind(&userInfo)
	util.Check(err)

	if len(userInfo.Email) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请输入有效邮箱！！"})
		return
	}

	sendmail.DayMail(userInfo.Email)

}

// UpdateAccountName 修改用户名称
func UpdateAccountName(c *gin.Context) {
	var userInfo util.UserInfo
	err := c.ShouldBind(&userInfo)
	util.Check(err)

	if len(userInfo.UserName) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账号不能为空"})
		return
	}

	if len(userInfo.Uname) == 0 {
		userInfo.Uname = util.RandomString(10)
	}

	if model.UpdateAccountName(userInfo.UserName, userInfo.Uname) == true {
		c.JSON(200, gin.H{
			"msg": "修改成功",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "修改名称失败"})
	}
}

// ToUpdate 修改密码
func ToUpdate(c *gin.Context) {
	c.HTML(200, "updatapwd/index.html", nil)
}

func DoUpdate(c *gin.Context) {
	var userInfo util.UserInfo
	err := c.ShouldBind(&userInfo)
	util.Check(err)

	if len(userInfo.PassWord) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "密码需要6位数"})
		return
	}

	if len(userInfo.Email) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "邮箱不能为空"})
		return
	}

	if model.Update(userInfo.PassWord, userInfo.Email) == false {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 421, "msg": "修改失败"})
		return
	}

	c.JSON(200, gin.H{
		"msg": "修改成功",
	})
}

// Info 用户信息路由
func Info(c *gin.Context) {

	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}














