package router

import (
	"task_manager/controllers"
	"task_manager/db"

	"github.com/gin-gonic/gin"
)

func RouterSetup(colls *db.DBCollections) *gin.Engine{ 
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	// Getting all the tasks
	router.GET("/tasks", controllers.GetTasks(colls))

	// Getting specific tasks
	router.GET("/tasks/:id", controllers.GetTask(colls))

	// Update a specific task
	router.PUT("/tasks/:id", controllers.UpdateTask(colls))

	// Deleted a specific task
	router.DELETE("/tasks/:id", controllers.RemoveTask(colls))

	// Post a specific task
	router.POST("/tasks", controllers.AddTask(colls))

	return router

}