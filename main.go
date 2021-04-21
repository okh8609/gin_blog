package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "pong"})
	})
	r.Run()
}
