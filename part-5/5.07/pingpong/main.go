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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	createTable()

	insertTableSQL := `INSERT INTO ping_pong (num) VALUES ($1);`

	r.GET("/pingpong", func(c *gin.Context) {
		db, err := connectDB()
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		defer db.Close()
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
	// for ingress check health
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/healthz", func(c *gin.Context) {
		db, err := connectDB()
		if err != nil {
			c.String(http.StatusInternalServerError, "Could not connect to the database")
			return
		}
		defer db.Close()
		c.String(http.StatusOK, "pong")
	})
	r.Run(":" + port)
}
func createTable() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
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
}
func connectDB() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	//fmt.Println(host, user, password, database)
	connStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", host, user, password, database)
	db, err := sql.Open("postgres", connStr)
	return db, err
}
