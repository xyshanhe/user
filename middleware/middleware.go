package middleware

import (
	"User/model/usermodel"
	"User/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //导入mysql
	"github.com/jinzhu/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := util.ParesToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/project?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		//验证通过后获取claim中的userid
		userId := claims.Userid
		var user usermodel.User
		db.First(&user, userId)

		//用户
		if user.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//用户存在 将user的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
