package usecases_test

import (
	"context"
	"errors"
	domain "task_management_clean_architecture/Domain"
	usecases "task_management_clean_architecture/Usecases"
	repositorymocks "task_management_clean_architecture/mocks/repository_mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	mockRepo *repositorymocks.UserRepositoryMock
	usecase domain.UserUsecase
	ctx context.Context
}

func (s *UserUsecaseTestSuite) SetupTest(){
	s.mockRepo = new(repositorymocks.UserRepositoryMock)
	s.usecase = usecases.NewUserUsecase(s.mockRepo, time.Second*2)
	s.ctx = context.Background()
}

func (s *UserUsecaseTestSuite) TestSignup(){
	s.Run("success", func ()  {
		newUser := &domain.User{Username: "newuser", Password: "pass123"}

		s.mockRepo.On("FindByUsername", mock.Anything, "newuser").Return(domain.User{}, errors.New("not found")).Once()

		s.mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(user *domain.User) bool{
			return user.Username == "newuser" && user.Password != "pass123"
		})).Return(nil).Once()

		err := s.usecase.Signup(s.ctx, newUser)
		s.NoError(err)
	})

	s.Run("UsernameAlreadyExists", func(){
		existingUser := domain.User{Username: "existinguser", Password: "somehash"}

		s.mockRepo.On("FindByUsername", mock.Anything, "existinguser").Return(existingUser, nil).Once()

		newUser := &domain.User{Username: "existinguser", Password: "pass123"}
		err := s.usecase.Signup(s.ctx, newUser)
		s.Error(err)
		s.EqualError(err, "username already exists")
	})
}

func (s *UserUsecaseTestSuite) TestLogin(){
	rawPassword := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	s.Require().NoError(err)

	mockUser := domain.User{
		UserID: primitive.NewObjectID(),
		Username: "user",
		Password: string(hashedPassword),
		Role: "admin",
	}

	s.mockRepo.On("FindByUsername", mock.Anything, "user").Return(mockUser, nil)

	user, token, err := s.usecase.Login(s.ctx, domain.User{
		Username: "user",
		Password: rawPassword,
	})

	s.NoError(err)
	s.Equal("user", user.Username)
	s.NotEmpty(token)
	s.mockRepo.AssertExpectations(s.T())
}

func TestUserUsecaseTestSuite(t *testing.T){
	suite.Run(t, new(UserUsecaseTestSuite))
}
