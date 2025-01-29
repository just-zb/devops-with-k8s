package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func generate() string {
	// generate the timestamp

	u := uuid.New()
	return fmt.Sprintf(u.String())

}
func main() {

	id := generate()
	ts := time.Now().UTC().Format(time.RFC3339)

	message := os.Getenv("MESSAGE")
	file, err := os.ReadFile("/config/information.txt")
	log.Println(string(file))
	log.Println("message:", message)
	if err != nil {
		log.Fatal(err)
	}
	buffer := bytes.NewBuffer(file)
	inf := buffer.String()

	r := gin.Default()
	r.GET("/log-output", func(c *gin.Context) {
		resp, err := http.Get("http://" + "pingpong-svc" + "/pingpong")
		if err != nil {
			log.Fatal(err)
		}
		var buffer bytes.Buffer
		log.Println("response body:", resp.Body)

		_, err = io.Copy(&buffer, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		message := "file content: " + inf + "\n" +
			"env variable: MESSAGE=" + message + "\n"
		c.String(http.StatusOK, message+ts+": "+id+"\n"+buffer.String())
	})
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.Run(":8080")
}
