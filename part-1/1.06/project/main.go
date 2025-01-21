package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
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

//go:embed templates/*
var templates embed.FS

func startServer(port string) {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/*")
	tmpl := template.Must(template.ParseFS(templates, "templates/*"))
	r.SetHTMLTemplate(tmpl)
	r.GET("/", func(c *gin.Context) {
		//c.String(http.StatusOK, "Hello World")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Hello devops",
			"message": "Hello golang",
		})
	})
	fmt.Printf("Server started in port %s\n", port)
	r.Run(":" + port)
}
