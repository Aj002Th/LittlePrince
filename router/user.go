package router

import (
	"encoding/gob"

	"github.com/Aj002Th/LittlePrince/db"
	"github.com/Aj002Th/LittlePrince/pkg/e"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(&db.User{})
}

type UserController struct {
}

func initUserController() *UserController {
	return new(UserController)
}

type LoginJson struct {
	Uname string `json:"uname"`
	Pwd   string `json:"pwd"`
}

// get session id by login -> /login
func (*UserController) Login(ctx *gin.Context) {
	loginJson := LoginJson{}
	err := ctx.Bind(&loginJson)
	if err != nil {
		ctx.JSON(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS))
		return
	}

	user, err := db.UserR.GetByAccount(loginJson.Uname, loginJson.Pwd)
	if err != nil {

		ctx.JSON(e.ERROR, e.GetMsg(e.ERROR))
		return
	}

	if err := login(ctx, *user); err != nil {
		ctx.JSON(e.ERROR, e.GetMsg(e.ERROR))
		return
	}

	ctx.JSON(e.SUCCESS, e.GetMsg(e.SUCCESS))
}

// clear session by logout -> /logout
func (*UserController) Logout(ctx *gin.Context) {
	if err := logout(ctx); err != nil {
		ctx.JSON(e.ERROR, e.GetMsg(e.ERROR))
		return
	}
}

// help function

// set user structure in session
func login(ctx *gin.Context, user db.User) error {
	session := sessions.Default(ctx)
	session.Set("user", user)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

// clear user structure in session
func logout(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	session.Clear()
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
