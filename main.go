package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v1",
		})
	})
	r.GET("/kill", func(c *gin.Context) {
		fmt.Println("Good Bye!!")
		os.Exit(1)
	})
	r.Run()
}
