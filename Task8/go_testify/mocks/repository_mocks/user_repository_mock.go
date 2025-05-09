package repositorymocks

import (
	"context"
	domain "task_management_clean_architecture/Domain"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(ctx context.Context, user *domain.User)error{
	arg := m.Called(ctx, user)
	return arg.Error(0)
}

func (m *UserRepositoryMock) FindByUsername(ctx context.Context, username string)(domain.User, error){
	args := m.Called(ctx, username)
	return args.Get(0).(domain.User), args.Error(1)
}