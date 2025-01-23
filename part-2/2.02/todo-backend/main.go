package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
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
	todos["Buy groceries"] = true
	todos["Clean the house"] = false
	todos["Pay bills"] = true
	r.GET("/todos", func(context *gin.Context) {
		context.JSON(http.StatusOK, todos)
	})
	r.POST("/todos", func(c *gin.Context) {
		var newTodo struct {
			Todo string `json:"todo"`
			Done bool   `json:"done"`
		}
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todos[newTodo.Todo] = newTodo.Done
		c.JSON(http.StatusOK, gin.H{"message": "Todo added"})
	})
	r.Run(":8080")
}
