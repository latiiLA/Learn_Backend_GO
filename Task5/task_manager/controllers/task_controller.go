package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTasks(taskCollection *mongo.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tasks, err := data.GetTasks(taskCollection)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch tasks from database",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
	}
}


func GetTask(taskCollection *mongo.Collection) gin.HandlerFunc{
	return func(ctx *gin.Context){
		// Get the task ID from URL parameters
		taskID := ctx.Param("id")

		// Call the service to fetch the task
		task, err := data.GetTask(taskCollection, taskID)
		if err != nil{
			if err == mongo.ErrNoDocuments{
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "Task not found",
				})
			}else{
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to fetch task",
				})
			}
			return
		}

		// Send the task and part of response
		ctx.JSON(http.StatusOK, gin.H{
			"task": task,
		})
	}
}

func AddTask(taskCollection *mongo.Collection) gin.HandlerFunc{
	return func(ctx *gin.Context){
		var newTask models.Task
		if err := ctx.ShouldBindJSON(&newTask); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if newTask.ID == ""{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
			return
		}

		if newTask.Title == ""{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
			return
		}

		if newTask.Description == ""{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
			return
		}

		if newTask.Status == ""{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
			return
		}

		// Call the service to add the task
		InsertedID, err := data.AddTask(taskCollection, newTask)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to add task",
			})
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Task created successfully", "task": newTask, "id": InsertedID})
	}
}

func UpdateTask(taskCollection *mongo.Collection) gin.HandlerFunc{
	return func(ctx *gin.Context){
		taskID := ctx.Param("id")
		if taskID == "" {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
            return
        }

		var updatedData models.Task
        if err := ctx.ShouldBindJSON(&updatedData); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid request body",
                "details": err.Error(),
            })
            return
        }

        update := bson.M{}
        if updatedData.Title != "" {
            update["title"] = updatedData.Title
        }
        if updatedData.Description != "" {
            update["description"] = updatedData.Description
        }
        if !updatedData.DueDate.IsZero() {
            update["due_date"] = updatedData.DueDate
        }
        if updatedData.Status != "" {
            update["status"] = updatedData.Status
        }

        if len(update) == 0 {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
            return
        }

		// Call the service to update the task
		MatchedCount, ModifiedCount, err := data.UpdateTask(taskCollection, taskID, update)
		if err != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to update task",
			})
			return
		}

	
		if MatchedCount == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		
		// Send response
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Task updated successfully",
			"updatedCount": ModifiedCount,
		})
	}
}

func RemoveTask(taskCollection *mongo.Collection) gin.HandlerFunc{
	return func(ctx *gin.Context){
		taskID := ctx.Param("id")
		if taskID == ""{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
			return
		}

		// Call the service to update the task
		DeletedCount, err := data.RemoveTask(taskCollection, taskID)
		print(DeletedCount, err)
		if err != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to remove task",
			})
			return
		}

		if DeletedCount == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		// Send response
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Task removed successfully",
			"deletedCount": DeletedCount,
		})
	}
}
