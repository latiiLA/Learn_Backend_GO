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

type TaskRepositoryTestSuite struct {
	suite.Suite
	DB *mongo.Database
	Collection string
	Repo domain.TaskRepository
	ctx context.Context
	Cleanup func()
}

func (s *TaskRepositoryTestSuite) SetupSuite(){
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOpts)
	s.Require().NoError(err)

	db := client.Database("test_task_suite")
	s.DB = db
	s.Collection = "tasks"
	s.ctx = context.TODO()

	s.Repo = repositories.NewTaskRepository(db, s.Collection)

	// cleanup function
	s.Cleanup = func() {
		_ = db.Drop(context.TODO())
		_ = client.Disconnect(context.TODO())
	}
}

func (s *TaskRepositoryTestSuite) TearDownTest(){
	// Clean up collection after each test
	err := s.DB.Collection(s.Collection).Drop(s.ctx)
	s.Require().NoError(err)
}

func (s *TaskRepositoryTestSuite) TestCreateTask(){
	task := &domain.Task{
		ID : primitive.NewObjectID(),
		UserID: primitive.NewObjectID(),
		Title: "Create Test",
		Status: "open",
	}

	err := s.Repo.Create(s.ctx, task)
	s.Require().NoError(err)

	// verify task was inserted
	count, err := s.DB.Collection(s.Collection).CountDocuments(s.ctx, bson.M{"_id": task.ID})
	s.Require().NoError(err)
	s.Require().Equal(int64(1), count)
}

func (s *TaskRepositoryTestSuite) TestFetchByUserID(){
	userID := primitive.NewObjectID()
	task := &domain.Task{
		ID: primitive.NewObjectID(),
		UserID : userID,
		Title: "User Task",
		Status: "in_progress",
	}

	err := s.Repo.Create(s.ctx, task)

	s.Require().NoError(err)

	tasks, err := s.Repo.FetchByUserID(s.ctx, userID)
	s.Require().NoError(err)
	s.Require().Len(tasks, 1)
	s.Require().Equal(userID, tasks[0].UserID)
}

func (s *TaskRepositoryTestSuite) TestUpdateByTaskID(){
	task := &domain.Task{
		ID: primitive.NewObjectID(),
		UserID : primitive.NewObjectID(),
		Title: "original",
		Status: "open",
	}
	err := s.Repo.Create(s.ctx, task)

	s.Require().NoError(err)

	taskID := task.ID

	update := &domain.Task{
		Title: "Updated Title",
		Status: "done",
	}

	matched, modified, err := s.Repo.UpdateByTaskID(s.ctx, taskID, update)
	s.Require().NoError(err)
	s.Require().EqualValues(1, matched)
	s.Require().EqualValues(1, modified)
}

func (s *TaskRepositoryTestSuite) TestDeleteByTaskID(){
	task := &domain.Task{
		ID: primitive.NewObjectID(),
		UserID : primitive.NewObjectID(),
		Title: "To be deleted",
		Status: "open",
	}
	err := s.Repo.Create(s.ctx, task)

	s.Require().NoError(err)

	taskID := task.ID

	deleteCount, err := s.Repo.DeleteByTaskID(s.ctx, taskID)
	s.Require().NoError(err)
	s.Require().EqualValues(1, deleteCount)
}

func (s *TaskRepositoryTestSuite) TestFetchByTaskID(){
	task := &domain.Task{
		ID: primitive.NewObjectID(),
		UserID : primitive.NewObjectID(),
		Title: "Find me",
		Status: "in_progress",
	}
	err := s.Repo.Create(s.ctx, task)

	s.Require().NoError(err)

	taskID := task.ID

	fetchedTask, err := s.Repo.FetchByTaskID(s.ctx, taskID)
	s.Require().NoError(err)
	s.Require().Equal(taskID, fetchedTask.ID)
	s.Require().Equal("Find me", fetchedTask.Title)	
}

func TestTaskRepositoryTestSuite(t *testing.T){
	suite.Run(t, new(TaskRepositoryTestSuite))
}