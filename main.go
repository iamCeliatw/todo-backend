package main

import (
	// "fmt"
	// "backend/controllers"
	"backend/controllers"
	"backend/database"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func main() {
	db := database.ConnectToDB()
	r := gin.Default()

	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos/:id", controllers.GetTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	r.Run(":8081")
	defer db.Close()
}
