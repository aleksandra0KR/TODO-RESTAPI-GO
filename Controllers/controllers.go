package controllers

import (
	"TODO-API-GO/entity"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var todos = []entity.Todo{
	{ID: "1", Item: "Clean room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Homework", Completed: false},
}

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func AddTodo(content *gin.Context) {
	var newTodo entity.Todo

	if err := content.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	content.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodoById(id string) (*entity.Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
func TakeTodoFromContext(content *gin.Context) *entity.Todo {
	id := content.Param("id")
	todo, err := GetTodoById(id)

	if err != nil {
		content.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not founded"})
		return nil
	}
	return todo
}
func GetTodo(content *gin.Context) {
	todo := TakeTodoFromContext(content)

	content.IndentedJSON(http.StatusOK, todo)
}

func ChangeStatus(content *gin.Context) {
	todo := TakeTodoFromContext(content)

	todo.Completed = !todo.Completed
	content.IndentedJSON(http.StatusOK, todo)
}
