package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aj002Th/LittlePrince/pkg/setting"
	"github.com/Aj002Th/LittlePrince/router"
	"github.com/Aj002Th/LittlePrince/service"
)

func main() {
	service.Setup()
	r := router.InitRouter()

	// go1.8+ 标准库内置了相关函数
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    setting.Server.HttpPort,
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

	// 旧版 go 使用 endless 实现优雅停机
	// server := endless.NewServer(setting.Server.HttpPort, r)
	// server.BeforeBegin = func(add string) {
	// 	pid := syscall.Getpid()
	// 	log.Printf("Actual pid is %d", pid)
	// 	ioutil.WriteFile("pid", []byte(fmt.Sprintf("%d", pid)), 0777)
	// }
	// if err := server.ListenAndServe(); err != nil {
	// 	log.Fatalln(err.Error())
	// }
}
