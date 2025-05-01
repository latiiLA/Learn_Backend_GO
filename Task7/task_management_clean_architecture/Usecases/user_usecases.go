package usecases

import (
	"context"
	"errors"
	"fmt"
	domain "task_management_clean_architecture/Domain"
	infrastructure "task_management_clean_architecture/Infrastructure"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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

func (uc *userUsecase) Signup(c context.Context, user *domain.User)error{
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	existingUser, err := uc.userRepository.FindByUsername(ctx, user.Username)
	if err == nil && existingUser.Username != "" {
		return errors.New("username already exists")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	user.Password = string(encryptedPassword)

	user.UserID = primitive.NewObjectID()
	
	return uc.userRepository.Create(ctx, user)
}

func (uc *userUsecase) Login(c context.Context, request domain.User)(domain.User, string, error){
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()
	
	user, err := uc.userRepository.FindByUsername(ctx, request.Username)
	if err != nil{
		return domain.User{}, "", errors.New("invalid username or password")
	}

	if !infrastructure.CheckPasswordHash(request.Password, user.Password){
		fmt.Println("Plain password from request:", request.Password)
		fmt.Println("Hashed password from DB:", user.Password)

		return domain.User{}, "", errors.New("invalid username or password1")
	}

	accessToken, err := infrastructure.GenerateToken(user.UserID, user.Role)
	if err != nil{
		return domain.User{}, "", err
	}

	return user, accessToken, nil
}