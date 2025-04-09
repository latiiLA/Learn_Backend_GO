package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/controllers"
)

func RouterSetup() *gin.Engine{ 
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	// Getting all the tasks
	router.GET("/tasks", controllers.GetTasks)

	// Getting specific tasks
	router.GET("/tasks/:id", controllers.GetTask)

	// Update a specific task
	router.PUT("/tasks/:id", controllers.UpdateTask)

	// Deleted a specific task
	router.DELETE("/tasks/:id", controllers.RemoveTask)

	// Post a specific task
	router.POST("/tasks", controllers.AddTask)

	return router

}