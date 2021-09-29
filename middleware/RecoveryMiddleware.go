package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "！!！"})
			}
		}()

		ctx.Next()
	}
}
