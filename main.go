package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"version": 1,
		})
	})

	r.GET("/time", func(c *gin.Context) {
		zone, offset := time.Now().Zone()
		c.JSON(http.StatusOK, gin.H{
			"zone": fmt.Sprintf("%s %d", zone, offset),
			"time": time.Now().Format("2006-01-02T15:04:05.000 MST"),
		})
	})

	fmt.Println("let's go!")
	r.Run(":3000")
}
