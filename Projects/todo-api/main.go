package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id        int    `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{Id: 1, Item: "Lavar a louça", Completed: false},
	{Id: 2, Item: "Academia", Completed: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var t Todo

	if err := c.BindJSON(&t); err != nil {
		return
	}

	todos = append(todos, t)
	c.IndentedJSON(http.StatusCreated, t)
}

func main() {

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run("localhost:8080")

}
