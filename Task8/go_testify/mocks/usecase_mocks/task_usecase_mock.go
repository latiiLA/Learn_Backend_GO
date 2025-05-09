package usecasemocks

import (
	"context"
	domain "task_management_clean_architecture/Domain"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseMock struct {
	mock.Mock
}

func (m *TaskUsecaseMock) AddTask(ctx context.Context, task *domain.Task) error{
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *TaskUsecaseMock) GetTasksByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Task, error){
	args := m.Called(ctx, userID)
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *TaskUsecaseMock) GetAllTasks(ctx context.Context)([]domain.Task, error){
	args := m.Called(ctx)
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *TaskUsecaseMock) UpdateTask(ctx context.Context, taskID primitive.ObjectID, update *domain.Task)(int64, int64, error){
	args := m.Called(ctx, taskID, update)
	return args.Get(0).(int64), args.Get(1).(int64), args.Error(2)
}

func (m *TaskUsecaseMock) DeleteTask(ctx context.Context, taskID primitive.ObjectID)(int64, error){
	args := m.Called(ctx, taskID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *TaskUsecaseMock) GetSpecificTask(ctx context.Context, taskID primitive.ObjectID)(domain.Task, error){
	args := m.Called(ctx, taskID)
	return args.Get(0).(domain.Task), args.Error(1)
}