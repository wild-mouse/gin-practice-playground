package sample

import "github.com/gin-gonic/gin"

func Hoge(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "foo",
	})
}
