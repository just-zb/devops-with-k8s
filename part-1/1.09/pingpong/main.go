package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	num := 0
	r.GET("/pingpong", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(num))
		num++
	})
	r.Run(":8081")
}
