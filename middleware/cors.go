package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 关闭同源策略
func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Max-Age", "172800")          // 缓存 option 请求信息 单位为秒
		ctx.Header("Access-Control-Allow-Credentials", "false") //  跨域请求是否需要带 cookie 信息 默认设置为true
		ctx.Set("content-type", "application/json")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}

		ctx.Next()
	}
}
