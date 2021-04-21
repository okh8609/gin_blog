package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/okh8609/gin_blog/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	tag := v1.NewTag()
	article := v1.NewArticle()
	engine_group_api_v1 := engine.Group("/api/v1")

	engine_group_api_v1.POST("/tags",tag.Create)
	engine_group_api_v1.DELETE("/tags/:id",tag.Delete)
	engine_group_api_v1.PUT("/tags/:id",tag.Update)
	engine_group_api_v1.PATCH("/tags/:id/state",tag.Update)
	engine_group_api_v1.GET("/tags",tag.List)

	engine_group_api_v1.POST("/articles",article.Create)
	engine_group_api_v1.DELETE("/articles/:id",article.Delete)
	engine_group_api_v1.PUT("/articles/:id",article.Update)
	engine_group_api_v1.PATCH("/articles/:id/state",article.Update)
	engine_group_api_v1.GET("/articles/:id",article.Get)
	engine_group_api_v1.GET("/articles",article.List)

	return engine
}
