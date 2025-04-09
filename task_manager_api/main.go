package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	// Getting all the tasks
	router.GET("/tasks", getTasks)

	// Getting specific tasks
	router.GET("/tasks/:id", getTask)

	// Update a specific task
	router.PUT("/tasks/:id", updateTask)

	router.DELETE("/tasks/:id", removeTask)

	router.POST("/tasks", addTask)

	router.Run()
}

func getTasks(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func getTask(ctx *gin.Context){
	id := ctx.Param("id")

	for _, task := range tasks{
		if task.ID == id{
			ctx.JSON(http.StatusOK, task)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func updateTask(ctx *gin.Context){
	id := ctx.Param("id")

	var UpdatedTask Task

	if err := ctx.ShouldBindJSON(&UpdatedTask); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks{
		if task.ID == id{
			// update only the specified fields
			if UpdatedTask.Title != ""{
				tasks[i].Title = UpdatedTask.Title
			}
			if UpdatedTask.Description != ""{
				tasks[i].Description = UpdatedTask.Description
			}
			ctx.JSON(http.StatusOK, gin.H{"messsage": "Task updated"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func removeTask(ctx *gin.Context){
	id := ctx.Param("id")

	for i, val := range tasks{
		if val.ID == id{
			tasks = append(tasks[:i], tasks[i + 1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task no found"})
}

func addTask(ctx *gin.Context){
	var newTask Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

// Task represents a task with its properties
type Task struct{
	ID 	string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate time.Time `json:"due_date"`
	Status string `json:"status"`
}

// Mock data for tasks
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}