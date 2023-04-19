package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID   string `json:"id"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

var todos = []Todo{
	{ID: "1", Item: "brush your teeth", Done: true},
	{ID: "2", Item: "do some excercise", Done: false},
	{ID: "3", Item: "take a bath", Done: true},
}

func getTodo(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addtodo(context *gin.Context) {
	var newtodo Todo
	if err := context.BindJSON(&newtodo); err != nil {
		return
	}
	todos = append(todos, newtodo)
	context.IndentedJSON(http.StatusCreated, newtodo)
}
func gettodoid(id string) (*Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}

	}
	return nil, errors.New("no todo found")
}

func gettodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := gettodoid(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func updatetodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := gettodoid(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	todo.Done = !todo.Done
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodo)
	router.GET("/todos/:id", gettodo)
	router.PATCH("/todos/:id", gettodo)
	router.POST("/todos/add", addtodo)
	router.Run("localhost:8080")
}
