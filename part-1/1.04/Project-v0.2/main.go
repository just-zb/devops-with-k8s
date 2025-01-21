package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	startServer(port)
}
func startServer(port string) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	fmt.Printf("Server started in port %s\n", port)
	r.Run(":" + port)
}
