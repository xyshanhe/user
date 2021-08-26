package reptile

import (
	"User/common"
	"User/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql" //导入mysql
)

var (
	// data = `<div class="link_item">[\s\S]+?<img[\s\S]+?src="([\s\S]+?)"[\s\S]+?<a[\s\S]+?href="([\s\S]+?)"[\s\S]+?>([\s\S]+?)<[\s\S]+?class="description a_animation">([\s\S]+?)<`
	data = `<div class="link_item">[\s\S]+?<img[\s\S]+?src="([\s\S]+?)"[\s\S]+?<a[\s\S]+?href="([\s\S]+?)"[\s\S]+?>([\s\S]+?)<[\s\S]+?class="description a_animation">([\s\S]+?)<[\s\S]+?class="category"[\s\S]+?>([\s\S]+?)<`
	Info []model.Data
)


//处理错误
func HandleErr(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func Fuun_data() {

	resp, err := http.Get("https://fuun.fun/page4/")
	HandleErr(err, `http.Get("https://fuun.fun/page4/")`)

	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)

	re := regexp.MustCompile(data)
	allString := re.FindAllStringSubmatch(html, -1)

	for _, x := range allString {
		Data_set(x[3], x[4], x[2], x[1], x[5])
	}
}

func Data_set(name, explain, href, images, category string) {

	common.DB.AutoMigrate(&model.Data{})

	// //增加

	common.DB.Create(&model.Data{Appname: name, Explain: explain, Addr: href, Imgs: images, Category: category})
}

func Data_get() {

	var d []model.Data

	common.Getslq()
	common.DB.Find(&d)

	//反转数据
	lengtg := len(d)
	for i:=0;i<lengtg/2;i++{
		temp := d[lengtg-1-i]
		d[lengtg-1-i] = d[i]
		d[i] = temp
	}

	Info = append(Info, d...)
}
