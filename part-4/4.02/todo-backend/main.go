package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var db *sql.DB
var dbLock sync.Mutex

func main() {
	go func() {
		for {
			if err := initDB(); err != nil {
				log.Println("init db failed", zap.Error(err))
				time.Sleep(5 * time.Second)
			} else {
				log.Println("init db success")
			}
		}
	}()
	startServer()
}
func startServer() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	})
	todos := make(map[string]bool)
	err := createTable(db)
	if err != nil {
		log.Fatal("create table failed", err)
	}

	r.GET("/todos", func(context *gin.Context) {
		todos, err = getTODO(db, todos)
		if err != nil {
			log.Fatal("get todos failed", err)
		}
		context.JSON(http.StatusOK, todos)
	})
	r.POST("/todos", func(c *gin.Context) {
		var newTodo struct {
			Todo string `json:"todo"`
			Done bool   `json:"done"`
		}
		logger, _ := zap.NewProduction()
		defer logger.Sync()

		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(newTodo.Todo) > 140 {
			logger.Warn("todo length is too long", zap.String("todo", newTodo.Todo))
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "todo too long",
			})
		} else {
			logger.Info("todo added", zap.String("todo", newTodo.Todo), zap.Bool("done", newTodo.Done))
			logger.Info("insert database")
			err := insertTODO(db, newTodo.Todo)
			if err != nil {
				log.Fatal("insert failed ", err)
			}
			c.JSON(http.StatusOK, gin.H{"message": "Todo added"})
		}

	})
	r.GET("/healthz", func(c *gin.Context) {
		if db != nil {
			c.JSON(http.StatusOK, gin.H{"healthy": true})
		} else {
			c.String(http.StatusInternalServerError, "DB is down")
		}
	})
	r.Run(":8080")
}
func createTable(db *sql.DB) error {
	createTableSQL := `create table if not exists todos(
    id serial primary key,
    todo text not null,
    done boolean not null default false
);`
	dbLock.Lock()
	defer dbLock.Unlock()
	_, err := db.Exec(createTableSQL)
	return err
}
func insertTODO(db *sql.DB, todo string) error {
	createTODOSQL := `insert into todos (todo) values ($1);`
	dbLock.Lock()
	defer dbLock.Unlock()
	_, err := db.Exec(createTODOSQL, todo)
	return err
}
func getTODO(db *sql.DB, todos map[string]bool) (map[string]bool, error) {
	getTODOSQL := `select * from todos`
	dbLock.Lock()
	defer dbLock.Unlock()
	rows, err := db.Query(getTODOSQL)
	if err != nil {
		return nil, err
	}
	type Todo struct {
		Id   int
		Todo string
		Done bool
	}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Todo, &todo.Done); err != nil {
			return nil, err
		}
		todos[todo.Todo] = todo.Done
	}
	return todos, nil
}
func initDB() error {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	connStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", host, user, password, database)
	tempDB, err := sql.Open("postgres", connStr)
	dbLock.Lock()
	defer dbLock.Unlock()
	db = tempDB
	return err
}
