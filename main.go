package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/internal/model"
	"github.com/okh8609/gin_blog/internal/routers"
	"github.com/okh8609/gin_blog/pkg/logger"
	"github.com/okh8609/gin_blog/pkg/setting"
	"github.com/okh8609/gin_blog/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// go build -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.getCommitID=`git rev-parse HEAD`"
	showVersion  bool
	buildTime    string
	buildVersion string
	getCommitID  string
)

var httpPort string
var runMode string
var configPathes string

func init() {
	setupFlag()

	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

// @title 部落格後端系統
// @version 1.0
// @description 再強一點：用Go語言完成六個大型專案
// @termsOfService https://github.com/okh8609/gin_blog
// @contact.name Khaos_Ou
// @host localhost:8080
func main() {
	if showVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", getCommitID)
		return
	}

	global.MyLogger.Infof(context.Background(), "%s: okh8609/%s", "Khaos", "gin_blog")
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
	ss, err := setting.NewSetting(strings.Split(configPathes, ",")...)
	if err != nil {
		return err
	}

	err = ss.ReadSection("Server", &global.Server)
	if err != nil {
		return err
	}

	err = ss.ReadSection("App", &global.App)
	if err != nil {
		return err
	}

	err = ss.ReadSection("Database", &global.Database)
	if err != nil {
		return err
	}

	err = ss.ReadSection("Auth", &global.Auth)
	if err != nil {
		return err
	}

	err = ss.ReadSection("Email", &global.Email)
	if err != nil {
		return err
	}

	if httpPort != "" {
		global.Server.HttpPort = httpPort
	}
	if runMode != "" {
		global.Server.RunMode = runMode
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.Database)
	return err
}

func setupLogger() error {
	path := global.App.LogSavePath + "/" + global.App.LogFileName + global.App.LogFileExt

	global.MyLogger = logger.NewLogger(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    500, // MB
		MaxAge:     28,  // Days
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false, // disabled by default
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("gin_blog", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = &jaegerTracer
	return nil
}

func setupFlag() error {
	flag.StringVar(&httpPort, "port", "8080", "啟動的通訊埠")
	flag.StringVar(&runMode, "mode", "debug", "執行的模式 [debug | release]")
	flag.StringVar(&configPathes, "cpath", "./configs/", "設定檔的尋找路徑")
	flag.BoolVar(&showVersion, "version", false, "編譯資訊")
	flag.Parse()
	return nil
}
