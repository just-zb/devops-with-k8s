package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
)

var uu string
var mu sync.Mutex

func main() {
	go uuidTest()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//mu.Lock()
		//defer mu.Unlock()
		c.String(http.StatusOK, uu)
	})
	r.Run(":8080")
}
func uuidTest() {
	u := uuid.New()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			//mu.Lock()
			uu = t.UTC().Format(time.RFC3339Nano) + ":" + u.String()
			//mu.Unlock()
		}
	}
}
