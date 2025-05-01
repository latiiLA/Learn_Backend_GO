package routers

import (
	"task_management_clean_architecture/Delivery/controllers"
	domain "task_management_clean_architecture/Domain"
	infrastructure "task_management_clean_architecture/Infrastructure"
	repositories "task_management_clean_architecture/Repositories"
	usecases "task_management_clean_architecture/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RouterSetup(jwtSecret string, timeout time.Duration, db *mongo.Database, router *gin.Engine){
	publicRouter := router.Group("")
	// All Public APIs
	NewUserRouter(timeout, db, publicRouter)
	// NewRefreshTokenRouter(timeout, db, publicRouter)

	protectedRouter := router.Group("")
	protectedRouter.Use(infrastructure.JwtAuthMiddleware(jwtSecret))
	NewTaskRouter(timeout, db, protectedRouter)
}

func NewUserRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup){
	ur := repositories.NewUserRepository(db, domain.CollectionUser)
	lc := &controllers.UserController{
		UserUsecase: usecases.NewUserUsecase(ur, timeout),
	}
	group.POST("/login", lc.Login)
	group.POST("/register", lc.Signup)
}

func NewTaskRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup){
	tr := repositories.NewTaskRepository(db, domain.CollectionTask)
	tc := &controllers.TaskController{
		TaskUsecase: usecases.NewTaskUsecase(tr, timeout),
	}
	group.GET("/tasks", tc.Fetch)
	group.POST("/tasks", tc.Create)
	group.GET("/alltasks", tc.FetchAll)
	group.PUT("/tasks/:id", tc.UpdateTask)
	group.DELETE("/tasks/:id", tc.DeleteTask)
	group.GET("/tasks/:id", tc.FetchSpecificTask)
}