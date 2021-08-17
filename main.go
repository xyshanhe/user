package main

import (
	"User/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	err := common.Getslq()
	if err != nil {
		panic(err)
	}

	defer common.DB.Close()

	//创建路由
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) //127.0.0.1:8080
}

