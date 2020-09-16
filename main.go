package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// get hostname
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		// get ip
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			panic(err)
		}

		var ip string
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
					break
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"hostname": hostname,
			"version":  "1.0.0",
			"ip":       ip,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
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
