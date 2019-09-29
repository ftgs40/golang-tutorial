package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		text := controllers.sayhello()
		c.JSON(200, gin.H{
			"message": text,
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
