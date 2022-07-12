package router

import (
	"io"
	"os"

	"github.com/Aj002Th/LittlePrince/middleware"

	"github.com/Aj002Th/LittlePrince/pkg/logging"
	"github.com/Aj002Th/LittlePrince/pkg/setting"
	v1 "github.com/Aj002Th/LittlePrince/router/api/v1"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var store redis.Store

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	// io
	f, err := logging.GetLogFile("router_")
	if err != nil {
		logging.Fatal("router.InitRouter err: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())

	// controller
	userCv1 := v1.InitUserController()

	// init session db : redis
	store, err = redis.NewStore(30, "tcp", setting.Redis.Host, setting.Redis.Password, []byte(setting.App.SessionKey))
	if err != nil {
		logging.Fatal("router.InitRouter err: %v", err)
	}
	r.Use(sessions.Sessions(setting.App.SessionName, store))

	r.POST("/register", userCv1.Register)
	r.POST("/login", userCv1.Login)
	r.DELETE("/logout", userCv1.Logout)

	api := r.Group("/api")
	api.Use(middleware.Authentication())
	{
		api.GET("/userinfo", userCv1.UserInfo)
	}

	return r
}
