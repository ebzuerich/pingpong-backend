package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"time": time.Now(),
		})
	})
	r.Run()
}