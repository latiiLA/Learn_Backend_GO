package router

import (
	"go-auth-task-manager/controllers"
	"go-auth-task-manager/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Accept TaskController instead of raw client
func RouterSetup(taskCollection *mongo.Collection, userCollection *mongo.Collection) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	// Routes
	router.GET("/tasks", middleware.AuthMiddleware(), middleware.AuthorizeRole("user"), controllers.GetTasks(taskCollection))
	router.GET("/tasks/:id", middleware.AuthMiddleware(), middleware.AuthorizeRole("user"), controllers.GetTask(taskCollection))
	router.PUT("/tasks/:id", middleware.AuthMiddleware(), middleware.AuthorizeRole("admin"), controllers.UpdateTask(taskCollection))
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.AuthorizeRole("admin"), controllers.RemoveTask(taskCollection))
	router.POST("/tasks", middleware.AuthMiddleware(), middleware.AuthorizeRole("admin"), controllers.AddTask(taskCollection))
	router.POST("/register", controllers.Register(userCollection))
	router.POST("/login", controllers.Login(userCollection))

	return router
}
