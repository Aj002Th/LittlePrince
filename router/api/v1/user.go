package v1

import (
	"encoding/gob"

	"github.com/Aj002Th/LittlePrince/db"
	"github.com/Aj002Th/LittlePrince/model"
	"github.com/Aj002Th/LittlePrince/model/request"
	"github.com/Aj002Th/LittlePrince/model/response"
	"github.com/Aj002Th/LittlePrince/pkg/e"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(&model.User{})
}

type UserController struct {
}

func InitUserController() *UserController {
	return new(UserController)
}

// sign in -> /register
func (*UserController) Register(ctx *gin.Context) {
	registerJson := request.UserInfo{}
	err := ctx.Bind(&registerJson)
	if err != nil {
		response.Fail(e.INVALID_PARAMS, ctx)
		return
	}

	user := model.User{Name: registerJson.Name, Pwd: registerJson.Pwd}
	user.ID, err = db.InsertUser(user)
	if err != nil {
		if err == nil {
			response.Fail(e.ERROR_EXIST_USERNAME, ctx)
			return
		}

		response.Fail(e.ERROR_INSERT_USER, ctx)
		return
	}

	if err := login(ctx, user); err != nil {
		response.Fail(e.ERROR_AUTH_SAVE_SESSION, ctx)
		return
	}

	response.Ok(ctx)
}

// get session id by login -> /login
func (*UserController) Login(ctx *gin.Context) {
	loginJson := request.UserInfo{}
	err := ctx.Bind(&loginJson)
	if err != nil {
		response.Fail(e.INVALID_PARAMS, ctx)
		return
	}

	user, err := db.GetUserByName(loginJson.Name)
	if err != nil {
		response.Fail(e.ERROR_WORONG_PWD, ctx)
		return
	}

	if err := login(ctx, user); err != nil {
		response.Fail(e.ERROR_AUTH_SAVE_SESSION, ctx)
		return
	}

	response.Ok(ctx)
}

// clear session by logout -> /logout
func (*UserController) Logout(ctx *gin.Context) {
	if err := logout(ctx); err != nil {
		response.Fail(e.ERROR_AUTH_SAVE_SESSION, ctx)
		return
	}

	response.Ok(ctx)
}

// get user info -> /userinfo
func (*UserController) UserInfo(ctx *gin.Context) {
	user := ctx.Keys["user"]
	response.OkWithData(user, ctx)
}

// help function

// set user structure in session
func login(ctx *gin.Context, user model.User) error {
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
