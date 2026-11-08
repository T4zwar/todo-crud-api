package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"todo-crud-api/database"
	"todo-crud-api/models"

	"github.com/gin-gonic/gin"
)

// CREATE: Add a new Todo
func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	todo.Title = strings.TrimSpace(todo.Title)

	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	result, err := database.DB.Exec(
		"INSERT INTO todos (title, completed) VALUES (?, ?)",
		todo.Title, todo.Completed,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Todo"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get Todo ID"})
		return
	}

	todo.ID = int(id)
	c.JSON(http.StatusCreated, todo)
}

// READ: Get all Todos
func GetTodos(c *gin.Context) {
	rows, err := database.DB.Query(
		"SELECT id, title, completed FROM todos ORDER BY id",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve Todos"})
		return
	}
	defer rows.Close()

	todos := []models.Todo{}

	for rows.Next() {
		var todo models.Todo

		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read Todos"})
			return
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// READ: Get one Todo by ID
func GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo ID"})
		return
	}

	var todo models.Todo

	err = database.DB.QueryRow(
		"SELECT id, title, completed FROM todos WHERE id = ?",
		id,
	).Scan(&todo.ID, &todo.Title, &todo.Completed)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve Todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UPDATE: Update an existing Todo
func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo ID"})
		return
	}

	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	todo.Title = strings.TrimSpace(todo.Title)

	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	result, err := database.DB.Exec(
		"UPDATE todos SET title = ?, completed = ? WHERE id = ?",
		todo.Title, todo.Completed, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Todo"})
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not verify update"})
		return
	}

	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	todo.ID = id
	c.JSON(http.StatusOK, todo)
}

// DELETE: Delete a Todo
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo ID"})
		return
	}

	result, err := database.DB.Exec(
		"DELETE FROM todos WHERE id = ?",
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete Todo"})
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not verify deletion"})
		return
	}

	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
