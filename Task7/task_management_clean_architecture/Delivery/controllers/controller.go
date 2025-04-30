package controllers

import (
	"fmt"
	"net/http"
	domain "task_management_clean_architecture/Domain"
	infrastructure "task_management_clean_architecture/Infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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

	err = tc.TaskUsecase.Create(c, &task)
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

	tasks, err := u.TaskUsecase.FetchByUserID(c, userID)
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

	_, err = uc.UserUsecase.FindByUsername(c, user.Username)
	if err == nil{
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Username already taken"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	user.Password = string(encryptedPassword)

	user.UserID = primitive.NewObjectID()

	err = uc.UserUsecase.Create(c, &user)
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

	user, err := uc.UserUsecase.FindByUsername(c, request.Username)
	if err != nil{
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found with the given username"})
		return
	}

	if !infrastructure.CheckPasswordHash(request.Password, user.Password){
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := infrastructure.GenerateToken(user.UserID, user.Role)
	if err != nil{
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": accessToken, "user": user,})
}