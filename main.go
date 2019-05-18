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
	r.GET("/v2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v2",
		})
	})
	r.GET("/v3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v3",
		})
	})
	r.GET("/v4", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v4",
		})
	})
	r.GET("/v5", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v5",
		})
	})
	r.GET("/v6", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v6",
		})
	})
	r.GET("/v7", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "v7",
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
