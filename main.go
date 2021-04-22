package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/internal/routers"
	"github.com/okh8609/gin_blog/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	gin.SetMode(global.Server.RunMode)
	engin := routers.NewRouter()
	engin.HandleMethodNotAllowed = global.Server.HandleMethodNotAllowed
	server := http.Server{
		Addr:           ":" + global.Server.HttpPort,
		Handler:        engin,
		ReadTimeout:    global.Server.ReadTimeout * time.Second,
		WriteTimeout:   global.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.Server)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.App)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.Database)
	if err != nil {
		return err
	}

	// fmt.Println(global.Server)
	// fmt.Println(global.App)
	// fmt.Println(global.Database)

	return nil
}
