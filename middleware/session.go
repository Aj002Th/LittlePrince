package middleware

import (
	"github.com/Aj002Th/LittlePrince/db"
	"github.com/Aj002Th/LittlePrince/model"
	"github.com/Aj002Th/LittlePrince/model/response"
	"github.com/Aj002Th/LittlePrince/pkg/e"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// auth
func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 查看 session 中是否有认证信息
		session := sessions.Default(ctx)
		sessionUser := session.Get("user")
		if sessionUser == nil {
			response.Fail(e.ERROR_AUTH_NOT_EXIST_SESSION, ctx)
			ctx.Abort()
			return
		}

		// 检查 session 中的认证信息是否失效
		user, err := db.GetUser(sessionUser.(*model.User).ID)
		if err != nil {
			response.Fail(e.ERROR_AUTH_INVALID_SESSION, ctx)
			ctx.Abort()
			return
		}

		// write obj into context
		ctx.Keys["user"] = user
		ctx.Next()
	}
}
