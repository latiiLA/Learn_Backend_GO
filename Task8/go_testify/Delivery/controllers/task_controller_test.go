package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task_management_clean_architecture/Delivery/controllers"
	domain "task_management_clean_architecture/Domain"
	usecasemocks "task_management_clean_architecture/mocks/usecase_mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerTestSuite struct {
	suite.Suite
	router *gin.Engine
	MockUsecase *usecasemocks.TaskUsecaseMock
	ctrl controllers.TaskController
}

func (s *TaskControllerTestSuite) SetupSuite(){
	gin.SetMode(gin.TestMode)
	s.router = gin.Default()

	s.MockUsecase = new(usecasemocks.TaskUsecaseMock)
	s.ctrl = controllers.TaskController{TaskUsecase: s.MockUsecase}

	s.router.POST("/tasks", func(c *gin.Context){
		c.Set("userID", primitive.NewObjectID().Hex())
		s.ctrl.Create(c)
	})
}

func (s *TaskControllerTestSuite) TestCreate(){
	s.Run("Success", func() {
		task := domain.Task{
			Title: "Test Task",
			Status: "open",
		}
	
		jsonData, _ := json.Marshal(task)
	
		s.MockUsecase.On("AddTask", mock.Anything, mock.AnythingOfType("*domain.Task")).Return(nil).Once()
	
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("userID", primitive.NewObjectID().Hex())

		s.ctrl.Create(c)
	
		s.Equal(http.StatusOK, w.Code)
		s.Contains(w.Body.String(), "Task created successfully")
	
		s.MockUsecase.AssertExpectations(s.T())
	})

	s.Run("Failure_InvalidJSON", func ()  {
		invalidJSON := []byte(`{"Title": "Invalid JSON"`)

		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(invalidJSON))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		
		s.ctrl.Create(c)

		s.Equal(http.StatusBadRequest, w.Code)
		s.Contains(w.Body.String(), "unexpected EOF")

		s.MockUsecase.AssertNotCalled(s.T(), "AddTask")
	})
	
}

func (s *TaskControllerTestSuite) TestFetch(){
	userID := primitive.NewObjectID()
	expectedTasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1"},
		{ID: primitive.NewObjectID(), Title: "Task 2"},
	}

	s.MockUsecase.On("GetTasksByUserID", mock.Anything, userID).Return(expectedTasks, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("userID", userID.Hex())

	s.ctrl.Fetch(c)

	s.Equal(http.StatusOK, w.Code)
	s.MockUsecase.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestFetchAll(){
	expectedTasks := []domain.Task{
		{Title: "Task A"},
		{Title: "Task B"},
	}

	s.MockUsecase.On("GetAllTasks", mock.Anything).Return(expectedTasks, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	s.ctrl.FetchAll(c)

	s.Equal(http.StatusOK, w.Code)
	s.MockUsecase.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestUpdateTask(){
	userID := primitive.NewObjectID()
	taskID := primitive.NewObjectID()

	updatedData := domain.Task{
		Title: "Updated Task Title",
	}

	// Marshal the updated task
	jsonData, _ := json.Marshal(updatedData)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+taskID.Hex(), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: taskID.Hex()}}
	c.Set("userID", userID.Hex())

	s.MockUsecase.On("UpdateTask", mock.Anything, taskID, &updatedData).Return(int64(1), int64(2), nil)
	s.ctrl.UpdateTask(c)
	s.Equal(http.StatusOK, w.Code)
	s.MockUsecase.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestDeleteTask(){
	taskID := primitive.NewObjectID()
	req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+taskID.Hex(),bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ :=gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: taskID.Hex()}}

	s.MockUsecase.On("DeleteTask", mock.Anything, taskID).Return(int64(1), nil)
	s.ctrl.DeleteTask(c)
	s.Equal(http.StatusOK, w.Code)
	s.MockUsecase.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestFetchSpecificTask(){
	expectedTasks := domain.Task{
		Title: "Task A",
	}

	taskID := primitive.NewObjectID()
	req, _ := http.NewRequest(http.MethodGet, "/tasks/"+taskID.Hex(), bytes.NewBuffer([]byte{}))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: taskID.Hex()}}
	
	s.MockUsecase.On("GetSpecificTask", mock.Anything, taskID).Return(expectedTasks, nil)
	s.ctrl.FetchSpecificTask(c)
	
	s.Equal(http.StatusOK, w.Code)
	s.MockUsecase.AssertExpectations(s.T())
}

func TestTaskControllerTestSuite(t *testing.T){
	suite.Run(t, new(TaskControllerTestSuite))
}