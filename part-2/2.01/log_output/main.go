package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
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
	r := gin.Default()
	r.GET("/log-output", func(c *gin.Context) {
		resp, err := http.Get("http://" + "ping" + ":80" + "/pingpong")
		if err != nil {
			log.Fatal(err)
		}
		var buffer bytes.Buffer
		log.Println("response body:", resp.Body)

		_, err = io.Copy(&buffer, resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusOK, ts+": "+id+"\n"+buffer.String())
	})
	r.Run(":8080")
}
