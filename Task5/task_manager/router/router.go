package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Accept TaskController instead of raw client
func RouterSetup(taskCollection *mongo.Collection) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	// Routes
	router.GET("/tasks", controllers.GetTasks(taskCollection))
	router.GET("/tasks/:id", controllers.GetTask(taskCollection))
	router.PUT("/tasks/:id", controllers.UpdateTask(taskCollection))
	router.DELETE("/tasks/:id", controllers.RemoveTask(taskCollection))
	router.POST("/tasks", controllers.AddTask(taskCollection))

	return router
}
