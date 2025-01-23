package main

import (
	"crypto/tls"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
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
	r.StaticFile("/image.jpg", "/usr/src/app/files/image.jpg")
	r.GET("/", func(c *gin.Context) {
		////c.String(http.StatusOK, "Hello World")
		log.Println("get a request")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "todo-app",
			"message": "todo-app",
			"image":   "/image.jpg",
		})
	})
	fmt.Printf("Server started in port %s\n", port)
	fmt.Println("add a goroutine to add local image")
	go addImage()

	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
func getImage() io.ReadCloser {
	url := "https://picsum.photos/1200"
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 跳过 TLS 证书验证
			},
		},
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("failed to download image: %s", resp.Status)
	}

	fmt.Println("image got")
	return resp.Body
}
func addImage() {
	fmt.Println("adding image to volume")
	out, err := os.Create("/usr/src/app/files/image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	body := getImage()
	defer body.Close()
	_, err = io.Copy(out, body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
	tick := time.Tick(1 * time.Hour)
	for {
		select {
		case <-tick:
			out, err := os.Create("/usr/src/app/files/image.jpg")
			if err != nil {
				log.Fatal(err)
			}
			body := getImage()
			_, err = io.Copy(out, body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("done")
		}
	}
}
