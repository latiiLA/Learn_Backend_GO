package controllers

import (
	"net/http"
	"task_manager/db"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTasks(collections *db.DBCollections) func (ctx *gin.Context){
	return func (ctx *gin.Context) {

		cursor, err := collections.Tasks.Find(ctx, bson.M{})
		if err != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch tasks from database",
			})
			return
		}
		defer cursor.Close(ctx)

		var tasks []models.Task
		if err = cursor.All(ctx, &tasks); err != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode task data",})
			return
		}

		if tasks == nil{
			tasks = []models.Task{}
		}

		ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
	}
}

func GetTask(collections *db.DBCollections) gin.HandlerFunc{
	return func (ctx *gin.Context){
		taskID := ctx.Param("id")

		var task models.Task
		err := collections.Tasks.FindOne(ctx, bson.M{"id": taskID}).Decode(&task)
		if err != nil{
			if err == mongo.ErrNoDocuments{
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "Task not found",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task",})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"task" : task})
	}
}

func UpdateTask(collections *db.DBCollections) gin.HandlerFunc{
	return func (ctx *gin.Context){
		taskID := ctx.Param("id") // This is a string-based ID
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

		result, err := collections.Tasks.UpdateOne(
            ctx,
            bson.M{"id": taskID},
            bson.M{"$set": update},
        )

		if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to update task",
                "details": err.Error(),
            })
            return
        }

        if result.MatchedCount == 0 {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
            return
        }

        ctx.JSON(http.StatusOK, gin.H{
            "message": "Task updated successfully",
            "updatedCount": result.ModifiedCount,
        })
	}
}

func RemoveTask(collections *db.DBCollections) gin.HandlerFunc{
	return func (ctx *gin.Context){
		taskID := ctx.Param("id")
		if taskID == ""{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
			return
		}

		result, err := collections.Tasks.DeleteOne(ctx, bson.M{"id": taskID})
		if err != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"}) 
			return
		}

		if result.DeletedCount == 0{
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully", "deletedCount" : result.DeletedCount})
	}
}

func AddTask(collections *db.DBCollections) gin.HandlerFunc{
	return func (ctx *gin.Context){
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

		result, err := collections.Tasks.InsertOne(ctx, newTask)
		if err != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create task",
			})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"message": "Task created successfully", "task": newTask, "id": result.InsertedID})
	}
}