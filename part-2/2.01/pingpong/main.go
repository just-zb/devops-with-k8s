package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	num := 0

	r.GET("/pingpong", func(c *gin.Context) {
		num++
		log.Println("received pingpong from:", c.Request.RemoteAddr)
		c.String(http.StatusOK, "Ping / Pongs: "+fmt.Sprint(num))
	})
	r.Run(":8080")
}
