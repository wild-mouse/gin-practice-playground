package main

import "github.com/gin-gonic/gin"
import "github.com/wild-mouse/gin-practice-playground/pkg/sample"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hoge", sample.Hoge)

	r.Run()
}
