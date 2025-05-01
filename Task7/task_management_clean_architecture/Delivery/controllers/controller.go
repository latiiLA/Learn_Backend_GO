package controllers

import (
	"fmt"
	"net/http"
	domain "task_management_clean_architecture/Domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

type UserController struct{
	UserUsecase domain.UserUsecase
}

func (tc *TaskController) Create(c *gin.Context){
	var task domain.Task

	err := c.ShouldBind(&task)
	if err != nil{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("userID")
	task.ID = primitive.NewObjectID()

	task.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.TaskUsecase.AddTask(c, &task)
	if err != nil{
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Task created succcessfully"})
}

func (tc *TaskController) Fetch(c *gin.Context){
	userID := c.GetString("userID")
	if userID == "" {
        c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User ID not found in token"})
        return
    }

	fmt.Println("Extracted userID:", userID)

	tasks, err := tc.TaskUsecase.GetTasksByUserID(c, userID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) FetchAll(c *gin.Context){
	tasks, err := tc.TaskUsecase.GetAllTasks(c)
	if err != nil{
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) UpdateTask(c *gin.Context){
	// Get task ID from URL params and updated tasks Data from body
	taskID := c.Param("id")
	if taskID == ""{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Task ID is required"})
		return
	}

	var updatedData domain.Task

	err := c.ShouldBind(&updatedData)
	if err != nil{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	fmt.Print("Reached the controller", taskID, updatedData)

	taskObjID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	matchedCount, ModifiedCount, err := tc.TaskUsecase.UpdateTask(c, taskObjID, &updatedData)
	if err != nil{
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	if matchedCount == 0{
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "No task found with the given ID"})
		return
	}

	message := "Task updated successfully"
	if ModifiedCount == 0{
		message = "Task found but no fields were updated"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"matchedCount": matchedCount,
		"modifiedCount": ModifiedCount,
	})
}

func (tc *TaskController) DeleteTask(c *gin.Context){
	taskID := c.Param("id")
	if taskID == ""{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Task ID is required"})
		return
	}

	taskObjID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	DeletedCount, err := tc.TaskUsecase.DeleteTask(c, taskObjID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	message := "Task Deleted Successfully"
	if DeletedCount == 0{
		message = "No task found with the given task ID"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"deletedCount": DeletedCount,
	})

}

func (tc *TaskController) FetchSpecificTask(c *gin.Context){
	taskID := c.Param("id")
	if taskID == ""{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Task ID is required"})
		return
	}

	taskObjID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	task, err := tc.TaskUsecase.GetSpecificTask(c, taskObjID)
	if err != nil{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task fetched successfully",
		"task": task,
	})
}

func (uc *UserController) Signup(c *gin.Context){
	var user domain.User

	err := c.ShouldBind(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	err = uc.UserUsecase.Signup(c, &user)
	if err != nil{
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User created successfully",
	})
}

func (uc *UserController) Login(c *gin.Context){
	var request domain.User

	err := c.ShouldBind(&request)
	if err != nil{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	user, accessToken, err := uc.UserUsecase.Login(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": accessToken, "user": user,})
}