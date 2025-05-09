package usecasemocks

import (
	"context"
	domain "task_management_clean_architecture/Domain"

	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock

}

func (m *UserUsecaseMock) Signup(ctx context.Context, user *domain.User)error{
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *UserUsecaseMock) Login(ctx context.Context, request domain.User)(domain.User, string, error){
	args := m.Called(ctx, request)
	return args.Get(0).(domain.User), args.String(1), args.Error(2)
}