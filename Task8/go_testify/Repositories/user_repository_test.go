package repositories_test

import (
	"context"
	domain "task_management_clean_architecture/Domain"
	repositories "task_management_clean_architecture/Repositories"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	DB *mongo.Database
	Collection string
	Repo domain.UserRepository
	ctx context.Context
	Cleanup func()
}

func (s *UserRepositoryTestSuite) SetupSuite(){
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOpts)
	s.Require().NoError(err)

	db := client.Database("test_task_suite")
	s.DB = db
	s.Collection = "tasks"
	s.ctx = context.TODO()

	s.Repo = repositories.NewUserRepository(db, s.Collection)

	// cleanup function
	s.Cleanup = func() {
		_ = db.Drop(context.TODO())
		_ = client.Disconnect(context.TODO())
	}
}

func (s *UserRepositoryTestSuite) TearDownTest(){
	// Clean up collection after each test
	err := s.DB.Collection(s.Collection).Drop(s.ctx)
	s.Require().NoError(err)
}

func (s *UserRepositoryTestSuite) TestCreate(){
	user := &domain.User{
		UserID: primitive.NewObjectID(),
		Username: "test_user",
		Password: "secret123",
		Role: "admin",
	}

	err := s.Repo.Create(s.ctx, user)
	s.Require().NoError(err)

	count, err := s.DB.Collection(s.Collection).CountDocuments(s.ctx, bson.M{"username": "testuser_create"})
	s.Require().NoError(err)
	s.Require().Equal(int64(1), count)
}

func (s *UserRepositoryTestSuite) TestFindByUsername(){
	user := &domain.User{
		UserID: primitive.NewObjectID(),
		Username: "testuser_find",
		Password: "anotherpassword",
		Role: "admin",
	}

	err := s.Repo.Create(s.ctx, user)
	s.Require().NoError(err)

	foundUser, err := s.Repo.FindByUsername(s.ctx, user.Username)
	s.Require().NoError(err)
	s.Require().Equal(user.Username, foundUser.Username)
	s.Require().Equal(user.Password, foundUser.Password)
	s.Require().Equal(user.Role, foundUser.Role)
}


func TestUserRepositoryTestSuite(t *testing.T){
	suite.Run(t, new(TaskRepositoryTestSuite))
}