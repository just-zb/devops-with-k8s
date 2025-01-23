package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	num := 0
	db := connectDB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		log.Fatalf("Could not establish a connection: %v", err)
	}
	log.Println("create table")
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS ping_pong (
		id SERIAL PRIMARY KEY,
		num INT NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Could not create table: %v", err)
	}
	insertTableSQL := `INSERT INTO ping_pong (num) VALUES ($1);`

	r.GET("/pingpong", func(c *gin.Context) {
		num++

		log.Println("received pingpong from:", c.Request.RemoteAddr)
		c.String(http.StatusOK, "Ping / Pongs: "+fmt.Sprint(num))

		log.Println("insert into table ping_pong, num:", num)
		_, err = db.Exec(insertTableSQL, num)
		if err != nil {
			log.Fatalf("Could not insert into table: %v", err)
		}
		log.Println("inserted")
	})
	r.Run(":8080")
}
func connectDB() *sql.DB {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	fmt.Println(host, user, password, database)
	connStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", host, user, password, database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return db
}
