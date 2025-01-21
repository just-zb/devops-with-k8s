package main

import (
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"time"
)

func generate() {
	// generate the timestamp
	file, err := os.OpenFile("/usr/src/app/files/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	tick := time.Tick(time.Second * 5)
	for {
		select {
		case <-tick:
			ts := time.Now().UnixMilli()
			if _, err := file.WriteString(fmt.Sprint(ts) + "\n"); err != nil {
				log.Fatal(err)
			}
			fmt.Println("timestamp appended", ts)
		}
	}

}
func getHash() uint32 {

	file, err := os.ReadFile("/usr/src/app/files/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	h := fnv.New32()
	h.Write(file)
	return h.Sum32()
}
func main() {
	if os.Getenv("GENERATE") == "true" {
		generate()
	} else {
		os.Create("/usr/src/app/files/log.txt")
		tick := time.Tick(time.Second * 10)
		for {
			select {
			case <-tick:
				hash := getHash()
				fmt.Println("hash value:", int(hash))
			}
		}
	}

}
