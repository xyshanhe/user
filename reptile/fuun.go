package reptile

import (
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
	Info []Data
)

type Data struct {
	Id       int
	Appname  string
	Explain  string
	Addr     string
	Imgs     string
	Category string
}

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

	model.DB.AutoMigrate(&Data{})

	// //增加

	model.DB.Create(&Data{Appname: name, Explain: explain, Addr: href, Imgs: images, Category: category})
}

func Data_get() {

	var d []Data

	model.DB.Find(&d)

	Info = append(Info, d...)
}
