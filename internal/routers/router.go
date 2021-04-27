package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/internal/middleware"
	"github.com/okh8609/gin_blog/internal/routers/api"
	v1 "github.com/okh8609/gin_blog/internal/routers/api/v1"
	"github.com/okh8609/gin_blog/pkg/limiter"
	"github.com/okh8609/gin_blog/pkg/upload"

	_ "github.com/okh8609/gin_blog/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var myLimiter = limiter.NewMyLimiter().AddBuckets(
	limiter.BucketInfo{
		Name:      "/auth",
		Capacity:  50,
		Frequency: 2 * time.Second,
		Quantum:   50,
	}, limiter.BucketInfo{
		Name:      "/auth/verify",
		Capacity:  50,
		Frequency: 2 * time.Second,
		Quantum:   50,
	})

func NewRouter() *gin.Engine {
	engine := gin.New()
	if global.Server.RunMode == "debug" {
		engine.Use(gin.Logger(), gin.Recovery())
	} else {
		engine.Use(middleware.AccessLog, middleware.Recovery)
	}
	engine.Use(middleware.RateLimiter(myLimiter))
	engine.Use(middleware.ContextTimeout(global.App.DefaultContextTimeout * time.Second))
	// engine.Use(middleware.TranslationMiddleware)

	url := ginSwagger.URL("http://kh-vm20:8080/swagger/doc.json") // The url pointing to API definition
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	tag := v1.NewTag()
	article := v1.NewArticle()
	engine_group_api_v1 := engine.Group("/api/v1")
	engine_group_api_v1.Use(middleware.JWTMiddleware)

	engine_group_api_v1.POST("/tags", tag.Create)
	engine_group_api_v1.DELETE("/tags/:id", tag.Delete)
	engine_group_api_v1.PUT("/tags/:id", tag.Update)
	engine_group_api_v1.PATCH("/tags/:id/state", tag.Update)
	engine_group_api_v1.GET("/tags", tag.List)

	engine_group_api_v1.POST("/articles", article.Create)
	engine_group_api_v1.DELETE("/articles/:id", article.Delete)
	engine_group_api_v1.PUT("/articles/:id", article.Update)
	engine_group_api_v1.PATCH("/articles/:id/state", article.Update)
	engine_group_api_v1.GET("/articles/:id", article.Get)
	engine_group_api_v1.GET("/articles", article.List)

	engine.POST("/upload/file", api.UploadFile)
	engine.StaticFS(upload.GetServerUrl(), http.Dir(global.App.UploadSavePath))

	engine.POST("/auth", api.CreateAuth)
	engine.POST("/auth/verify", api.VerifyAuth)
	engine.PUT("/auth", api.UpdateAuth)
	engine.DELETE("/auth", api.DeleteAuth)

	engine.GET("panic", func(c *gin.Context) { panic("??") })

	return engine
}
