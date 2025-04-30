package usecases

import (
	"context"
	domain "task_management_clean_architecture/Domain"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase{
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uc *userUsecase) Create(c context.Context, user *domain.User)error{
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.Create(ctx, user)
}

func (uc *userUsecase) FindByUsername(c context.Context, username string)(domain.User, error){
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	
	user, err := uc.userRepository.FindByUsername(ctx, username)
	if err != nil{
		return domain.User{}, err
	}

	return user, nil
}