package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Read book", Completed: true},
	{ID: "2", Item: "Clean room", Completed: true},
	{ID: "3", Item: "Call Sara", Completed: true},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newToduo todo

	if err := context.BindJSON(&newToduo); err != nil {
		return
	}

	todos = append(todos, newToduo)
	context.IndentedJSON(http.StatusCreated, todos)

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
