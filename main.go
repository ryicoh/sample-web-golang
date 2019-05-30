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
			"stage": "stg",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/v6", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v6",
		})
	})
	r.GET("/host", func(c *gin.Context) {
		hostname, err := os.Hostname()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
		}
		c.JSON(200, gin.H{
			"message": hostname,
		})
	})
	r.GET("/kill", func(c *gin.Context) {
		fmt.Println("Good Bye!!")
		os.Exit(1)
	})
	r.Run()
}
