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

func (tc *TaskController)Create(c *gin.Context){
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

func (u *TaskController) Fetch(c *gin.Context){
	userID := c.GetString("userID")
	if userID == "" {
        c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User ID not found in token"})
        return
    }

	fmt.Println("Extracted userID:", userID)

	tasks, err := u.TaskUsecase.GetTasksByUserID(c, userID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
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