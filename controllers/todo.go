package controllers

import (
	"backend/database"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todoList []models.Todo

	// 查詢數據庫
	rows, err := database.DB.Query("SELECT * FROM todoList")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// 遍歷每一行
	for rows.Next() {
		var todo models.Todo

		// 解析每一行數據到 todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 將 todo 添加到 todoList
		todoList = append(todoList, todo)
	}

	c.JSON(http.StatusOK, gin.H{"todoList": todoList})
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo

	// 解析請求體
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newTodo, "newTodo")
	// 插入數據到數據庫
	query := `INSERT INTO todoList (title, description) VALUES ($1, $2)`
	_, err := database.DB.Exec(query, newTodo.Title, newTodo.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo item created successfully."})
}

func GetTodo(c *gin.Context) {
	var todo models.Todo

	// 解析請求參數
	id := c.Param("id")

	// 查詢數據庫
	row := database.DB.QueryRow("SELECT * FROM todoList WHERE id=$1", id)

	// 解析數據到 todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func UpdateTodo(c *gin.Context) {
	var updatedTodo models.Todo

	// 解析請求參數
	id := c.Param("id")

	// 解析請求體
	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新數據庫
	result, err := database.DB.Exec("UPDATE todoList SET title=$1, description=$2, completed=$3 WHERE id=$4", updatedTodo.Title, updatedTodo.Description, updatedTodo.Completed, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 檢查是否有行被更新
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo item updated successfully."})
}

func DeleteTodo(c *gin.Context) {
	// 從 URL 中獲取 Todo 項目的 ID
	id := c.Param("id")
	// 執行刪除操作
	_, err := database.DB.Exec("DELETE FROM todoList WHERE id = $1", id)
	if err != nil {
		// 如果出現錯誤，返回錯誤響應
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回成功響應
	c.JSON(http.StatusOK, gin.H{"message": "Todo item deleted successfully."})
}
