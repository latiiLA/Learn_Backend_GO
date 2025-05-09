package controllers

import (
	"net/http"
	domain "task_management_clean_architecture/Domain"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	UserUsecase domain.UserUsecase
}

func (uc *UserController) Signup(c *gin.Context) {
	var user domain.User

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	err = uc.UserUsecase.Signup(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User created successfully",
	})
}

func (uc *UserController) Login(c *gin.Context) {
	var request domain.User

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	user, accessToken, err := uc.UserUsecase.Login(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": accessToken, "user": user})
}