package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks = []Task{
	{ID: 1, Name: "Learn Go"},
	{ID: 2, Name: "Build API with Gin"},
}

func main() {
	r := gin.Default()

	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, tasks)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var newTask Task
		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newTask.ID = len(tasks) + 1
		tasks = append(tasks, newTask)
		c.JSON(http.StatusCreated, newTask)
	})

	r.Run(":8080")
}
