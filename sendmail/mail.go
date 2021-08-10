package sendmail

/*
-----------------------------------
-----------------------------------
邮箱的配置信息。
-----------------------------------
-----------------------------------
*/

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

// Code ...
var (
	Code string
	//Data bool
)

func DayMail(mail string) {

	//获取随机6位数验证码
	rand.Seed(time.Now().Unix())
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	data := fmt.Sprint("【HiaHia】验证码是：", num)
	// 接收到的随机6位验证码给全局变量Code
	Code = fmt.Sprint(num)

	//发送邮件
	m := gomail.NewMessage()
	m.SetHeader("From", "2653563535@qq.com") //发送人
	m.SetHeader("To", mail)                  //收件人
	m.SetHeader("Subject", "【HiaHia分享】")     //邮件标题
	m.SetBody("text/html", data)             //邮件内容
	d := gomail.NewDialer("smtp.qq.com", 587, "2653563535@qq.com", "ntfqabvgdyjcdjeh")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
