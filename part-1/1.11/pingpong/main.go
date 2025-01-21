package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	r := gin.Default()
	num := 0
	file, err := os.Create("/usr/src/app/files/log.txt")
	file.Close()

	r.GET("/pingpong", func(c *gin.Context) {
		num++
		c.String(http.StatusOK, "pong "+fmt.Sprint(num))
		if err != nil {
			log.Fatal(err)
		}
		str := "Ping / Pongs: "
		err = os.WriteFile("/usr/src/app/files/log.txt", []byte(str+strconv.Itoa(num)), 0644)
		fmt.Println("Ping / Pongs: ", num)
	})
	r.Run(":8081")
}
