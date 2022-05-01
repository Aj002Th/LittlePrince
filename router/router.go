package router

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"github.com/Aj002Th/LittlePrince/middleware"
	"github.com/fvbock/endless"

	"github.com/Aj002Th/LittlePrince/pkg/logging"
	"github.com/Aj002Th/LittlePrince/pkg/setting"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var store redis.Store

func Setup() {
	r := InitRouter()
	server := endless.NewServer(setting.ServerSetting.HttpPort, r)
	server.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		log.Printf("Actual pid is %d", pid)
		ioutil.WriteFile("pid", []byte(fmt.Sprintf("%d", pid)), 0777)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err.Error())
	}
	r.Run(setting.ServerSetting.HttpPort)
}

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	// engine
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
	r.Use(middleware.Authentication())

	// controller
	userC := initUserController()

	// init session db : redis
	store, err = redis.NewStore(30, "tcp", setting.RedisSetting.Host, setting.RedisSetting.Password, []byte(setting.AppSetting.SessionSecret))
	if err != nil {
		logging.Fatal("router.InitRouter err: %v", err)
	}
	r.Use(sessions.Sessions(setting.AppSetting.SessionName, store))

	// log in & out
	r.POST("/login", userC.Login)
	r.DELETE("/logout", userC.Logout)

	// must login
	api := r.Group("/api")
	api.Use()
	{
		// business router
	}

	return r
}
