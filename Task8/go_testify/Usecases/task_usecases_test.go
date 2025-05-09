package usecases_test

import (
	"context"
	domain "task_management_clean_architecture/Domain"
	usecases "task_management_clean_architecture/Usecases"
	repositorymocks "task_management_clean_architecture/mocks/repository_mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseTestSuite struct {
	suite.Suite
	mockRepo *repositorymocks.TaskRepositoryMock
	usecase  domain.TaskUsecase
	ctx      context.Context
}

func (s *TaskUsecaseTestSuite) SetupTest() {
	s.mockRepo = new(repositorymocks.TaskRepositoryMock)
	s.usecase = usecases.NewTaskUsecase(s.mockRepo, time.Second*2)
	s.ctx = context.Background()
}


func (s *TaskUsecaseTestSuite) TestAddTask() {
	task := &domain.Task{Title: "Test Task"}

	s.mockRepo.On("Create", mock.Anything, task).Return(nil)

	err := s.usecase.AddTask(s.ctx, task)

	s.NoError(err)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseTestSuite) TestGetTasksByUserID(){
	expected := []domain.Task{{Title: "Task 1"}, {Title: "Task 2"}}
	userID := primitive.NewObjectID()

	s.mockRepo.On("FetchByUserID", mock.Anything, userID).Return(expected, nil)

	tasks, err := s.usecase.GetTasksByUserID(s.ctx, userID)

	s.NoError(err)
	s.Equal(expected, tasks)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseTestSuite) TestGetAllTasks(){
	expected := []domain.Task{{Title: "A"}, {Title: "B"}}

	s.mockRepo.On("FetchAll", mock.Anything).Return(expected, nil)

	tasks, err := s.usecase.GetAllTasks(s.ctx)

	s.NoError(err)
	s.Equal(expected, tasks)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseTestSuite) TestUpdateTask(){
	taskID := primitive.NewObjectID()
	updated := &domain.Task{Title: "updated"}

	s.mockRepo.On("UpdateByTaskID", mock.Anything, taskID, updated).Return(int64(1), int64(1), nil)

	modifiedCount, matchedCount, err := s.usecase.UpdateTask(s.ctx, taskID, updated)

	s.NoError(err)
	s.Equal(int64(1), modifiedCount)
	s.Equal(int64(1), matchedCount)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseTestSuite) TestDeleteTask(){
	taskID := primitive.NewObjectID()

	s.mockRepo.On("DeleteByTaskID", mock.Anything, taskID).Return(int64(1), nil)

	deletedCount, err := s.usecase.DeleteTask(s.ctx, taskID)

	s.NoError(err)
	s.Equal(int64(1), deletedCount)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseTestSuite) TestGetSpecificTask(){
	taskID := primitive.NewObjectID()
	expected := domain.Task{ID: taskID, Title: "Read Task"}

	s.mockRepo.On("FetchByTaskID", mock.Anything, taskID).Return(expected, nil)

	result, err := s.usecase.GetSpecificTask(s.ctx, taskID)

	s.NoError(err)
	s.Equal(expected, result)
	s.mockRepo.AssertExpectations(s.T())
}

func TestTaskUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseTestSuite))
}
