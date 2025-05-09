package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task_management_clean_architecture/Delivery/controllers"
	domain "task_management_clean_architecture/Domain"
	usecasemocks "task_management_clean_architecture/mocks/usecase_mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	router *gin.Engine
	MockUsecase *usecasemocks.UserUsecaseMock
	ctrl controllers.UserController
}

func (s *UserControllerTestSuite) SetupSuite(){
	gin.SetMode(gin.TestMode)
	s.router = gin.Default()

	s.MockUsecase = new(usecasemocks.UserUsecaseMock)
	s.ctrl = controllers.UserController{UserUsecase: s.MockUsecase}

	s.router.POST("/signup", s.ctrl.Signup)
	s.router.POST("/login", s.ctrl.Login)
}

func (s *UserControllerTestSuite) TestSignup(){
	s.Run("Success", func ()  {
		user := domain.User{
			Username: "testuser",
			Password: "testPass",
			Role: "user",
		}

		body, _ := json.Marshal(user)
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")
		
		s.MockUsecase.On("Signup", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		s.ctrl.Signup(c)
		s.Contains(w.Body.String(), "User created successfully")
		s.MockUsecase.AssertExpectations(s.T())
	})

	s.Run("BindingError", func ()  {
		invalidJSON := `{"username" : "testuser", "password":}`

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/signup", bytes.NewBufferString(invalidJSON))
		c.Request.Header.Set("Content-Type", "application/json")

		s.ctrl.Signup(c)

		s.Equal(http.StatusBadRequest, w.Code)
		s.Contains(w.Body.String(), "invalid character")
		s.MockUsecase.AssertNotCalled(s.T(), "Signup")
	})

	s.Run("UsecaseError", func ()  {
		user := domain.User{
			Username: "existinguser",
			Password: "testpass",
			Role: "user",
		}

		body, _ := json.Marshal(user)

		w := httptest.NewRecorder()
		c, _ :=gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		s.MockUsecase.On("Signup", mock.Anything, mock.AnythingOfType("*domain.User")).Return(errors.New("username already exists")).Once()
		
		s.ctrl.Signup(c)

		s.Equal(http.StatusInternalServerError, w.Code)
		s.Contains(w.Body.String(), "username already exists")
		s.MockUsecase.AssertExpectations(s.T())
	})
}

func (s *UserControllerTestSuite) TestLogin(){
	s.Run("Success", func() {
		request := domain.User{
			Username: "testuser",
			Password: "password",
		}

		expectedUser := domain.User{
			Username: "testuser",
			Role: "user",
		}

		expectedToken := "test_token_123"
	
		body, _ := json.Marshal(request)
		
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")
	
		s.MockUsecase.On("Login", mock.Anything, request).Return(expectedUser, expectedToken, nil).Once()
	
		s.ctrl.Login(c)
	
		s.Equal(http.StatusOK, w.Code)
		s.Contains(w.Body.String(), "Login successful")
		s.MockUsecase.AssertExpectations(s.T())
	})

	s.Run("InvalidJSON", func() {
		invalidJSON := `{"username": "testuser", "password":}`

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/login", bytes.NewBufferString(invalidJSON))
		c.Request.Header.Set("Content-Type", "application/json")

		s.ctrl.Login(c)

		s.Equal(http.StatusBadRequest, w.Code)
		s.Contains(w.Body.String(), "invalid character")
		s.MockUsecase.AssertNotCalled(s.T(), "Login")
	})

	s.Run("InvalidCredentials", func() {
		request := domain.User{
			Username: "wronguser",
			Password: "wrongpass",
		}

		body, _ := json.Marshal(request)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		s.MockUsecase.On("Login", mock.Anything, request).Return(domain.User{}, "", errors.New("invalid credentials")).Once()

		s.ctrl.Login(c)

		s.Equal(http.StatusInternalServerError, w.Code)
		s.Contains(w.Body.String(), "invalid credentials")
		s.MockUsecase.AssertExpectations(s.T())
	})
}

func TestUserControllerTestSuite(t *testing.T){
	suite.Run(t, new(UserControllerTestSuite))
}