package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"todo-crud-api/database"
	"todo-crud-api/handlers"
)

func main() {
	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Create Gin router
	router := gin.Default()

	// API routes
	router.POST("/todos", handlers.CreateTodo)
	router.GET("/todos", handlers.GetTodos)
	router.GET("/todos/:id", handlers.GetTodo)
	router.PUT("/todos/:id", handlers.UpdateTodo)
	router.DELETE("/todos/:id", handlers.DeleteTodo)

	// Start server
	log.Println("Server running on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
