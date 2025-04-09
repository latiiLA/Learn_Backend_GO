package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/data"
	"github.com/zaahidali/task_manager_api/models"
)

func GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"tasks": data.Tasks})
}

func GetTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, task := range data.Tasks {
		if task.ID == id {
			ctx.JSON(http.StatusOK, task)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var UpdatedTask models.Task

	if err := ctx.ShouldBindJSON(&UpdatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range data.Tasks {
		if task.ID == id {
			// update only the specified fields
			if UpdatedTask.Title != "" {
				data.Tasks[i].Title = UpdatedTask.Title
			}
			if UpdatedTask.Description != "" {
				data.Tasks[i].Description = UpdatedTask.Description
			}
			ctx.JSON(http.StatusOK, gin.H{"messsage": "Task updated"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func RemoveTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, val := range data.Tasks {
		if val.ID == id {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task no found"})
}

func AddTask(ctx *gin.Context) {
	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.Tasks = append(data.Tasks, newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}