package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palp1tate/brevinect/consts"
)

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, _ := ctx.Get("role")
		if role != consts.Admin {
			ctx.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
