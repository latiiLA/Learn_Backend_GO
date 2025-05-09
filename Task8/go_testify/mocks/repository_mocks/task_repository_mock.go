package repositorymocks

import (
	"context"
	domain "task_management_clean_architecture/Domain"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (m *TaskRepositoryMock) Create(ctx context.Context, task *domain.Task)error{
	arg := m.Called(ctx, task)
	return arg.Error(0)
}

func (m *TaskRepositoryMock) FetchByUserID(ctx context.Context, userID primitive.ObjectID)([]domain.Task, error){
	args := m.Called(ctx, userID)
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *TaskRepositoryMock) FetchAll(ctx context.Context)([]domain.Task, error){
	args := m.Called(ctx)
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *TaskRepositoryMock) UpdateByTaskID(ctx context.Context, taskID primitive.ObjectID, updatedData *domain.Task)(int64, int64, error){
	args := m.Called(ctx, taskID, updatedData)
	return args.Get(0).(int64), args.Get(1).(int64), args.Error(2)
}

func (m *TaskRepositoryMock) DeleteByTaskID(ctx context.Context, taskID primitive.ObjectID)(int64, error){
	args := m.Called(ctx, taskID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *TaskRepositoryMock) FetchByTaskID(ctx context.Context, taskID primitive.ObjectID)(domain.Task, error){
	args := m.Called(ctx, taskID)
	return args.Get(0).(domain.Task), args.Error(1)
}