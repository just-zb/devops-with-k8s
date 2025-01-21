package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	r := gin.Default()
	r.GET("/log-output", func(c *gin.Context) {
		file, err := os.ReadFile("/usr/src/app/files/log.txt")
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, ts+": "+id+"\n"+string(file))
	})
	r.Run(":8080")
}
