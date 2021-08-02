package reptile

import (
	"fmt"
	"regexp"
)

// https://www.friv.com/

// import (
// 	_ "github.com/go-sql-driver/mysql" //导入mysql
// )

var (
	data1    = `<div class="gameThumbContainer gameThumbLoading1"[\s\S]+?<a href="([\s\S]+?)">[\s\S]+?src="([\s\S]+?)"[\s\S]+?<div class="gameTitle notranslate" style="top: 3%;">([\s\S]+?)</div>`
	GameInfo []Gamedata
)

type Gamedata struct {
	Id       int
	Gamename string
	Addr     string
}

func Game_data() {

	// resp, err := http.Get("https://www.friv.com/")
	// HandleErr(err, `http.Get("https://www.friv.com/")`)

	// bytes, _ := ioutil.ReadAll(resp.Body)
	// html := string(bytes)

	re := regexp.MustCompile(data1)
	allString := re.FindAllStringSubmatch(D, -1)
	for _, x := range allString {
		d := "https://www.friv.com/"
		fmt.Println(d+x[1], d+x[2], x[3])
	}
}
