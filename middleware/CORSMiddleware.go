package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin","*")
		ctx.Writer.Header().Set("Access-Control-Max-Age","86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers","Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials","true")
		//ctx.Writer.Header().Set("Authorization","Bearer "+)

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		}else {
			ctx.Next()
		}
	}
}
