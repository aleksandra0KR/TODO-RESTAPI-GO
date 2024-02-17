package main

import (
	"TODO-API-GO/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", controllers.GetTodos)
	router.POST("/todos", controllers.AddTodo)
	router.GET("/todos/:id", controllers.GetTodo)
	router.PATCH("/todos/:id", controllers.ChangeStatus)
	err := router.Run("localhost:9090")
	if err != nil {
		return
	}
}
