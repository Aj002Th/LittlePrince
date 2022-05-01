package middleware

import (
	"github.com/Aj002Th/LittlePrince/db"
	"github.com/Aj002Th/LittlePrince/pkg/e"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// auth
func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get user from session
		session := sessions.Default(ctx)
		sessionUser := session.Get("user")
		if sessionUser == nil {
			ctx.JSON(e.UNAUTHORIZED, e.GetMsg(e.ERROR_AUTH))
			ctx.Abort()
			return
		}

		//check user in the db
		user, err := db.UserR.GetById(sessionUser.(*db.User).ID)
		if err != nil {
			ctx.JSON(e.UNAUTHORIZED, e.GetMsg(e.ERROR_AUTH))
			ctx.Abort()
			return
		}

		// write obj into context
		ctx.Keys["user"] = user
		ctx.Next()
	}
}
