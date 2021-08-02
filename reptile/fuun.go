package reptile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql" //导入mysql
	"github.com/jinzhu/gorm"
)

var (
	// data = `<div class="link_item">[\s\S]+?<img[\s\S]+?src="([\s\S]+?)"[\s\S]+?<a[\s\S]+?href="([\s\S]+?)"[\s\S]+?>([\s\S]+?)<[\s\S]+?class="description a_animation">([\s\S]+?)<`
	data = `<div class="link_item">[\s\S]+?<img[\s\S]+?src="([\s\S]+?)"[\s\S]+?<a[\s\S]+?href="([\s\S]+?)"[\s\S]+?>([\s\S]+?)<[\s\S]+?class="description a_animation">([\s\S]+?)<[\s\S]+?class="category"[\s\S]+?>([\s\S]+?)<`
	Info []Data
	sql  = "root:password@tcp(localhost:3306)/project?charset=utf8&parseTime=True&loc=Local"
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

	//用户名:密码@tcp(ip:port)/数据库名称?charset=utf8&paresTime=True&loc=Local
	db, err := gorm.Open("mysql", sql)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Data{})

	// //增加

	db.Create(&Data{Appname: name, Explain: explain, Addr: href, Imgs: images, Category: category})
}

func Data_get() {

	var d []Data

	//用户名:密码@tcp(ip:port)/数据库名称?charset=utf8&paresTime=True&loc=Local
	db, err := gorm.Open("mysql", sql)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Where("id >= ? ", 0).Find(&d)

	for _, j := range d {

		Info = append(Info, j)

	}
}
