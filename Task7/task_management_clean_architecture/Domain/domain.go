package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionTask = "tasks"
const CollectionUser = "users"

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"due_date"`
	Status      string `json:"status"`
	UserID primitive.ObjectID `bson:"userID" json:"-"`
}

type User struct {
	UserID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role string `json:"role" bson:"role"`
}

type TaskRepository interface{
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string)([]Task, error)
	FetchAll(c context.Context)([]Task, error)
	UpdateByTaskID(c context.Context, taskID primitive.ObjectID, updatedData *Task)(int64, int64, error)
	DeleteByTaskID(c context.Context, taskID primitive.ObjectID)(int64, error)
	FetchByTaskID(c context.Context, taskID primitive.ObjectID)(Task, error)
}

type TaskUsecase interface{
	AddTask(c context.Context, task *Task) error
	GetTasksByUserID(c context.Context, userID string)([]Task, error)
	GetAllTasks(c context.Context)([]Task, error)
	UpdateTask(c context.Context, taskID primitive.ObjectID, updatedData *Task)(int64, int64, error)
	DeleteTask(c context.Context, taskID primitive.ObjectID)(int64, error)
	GetSpecificTask(c context.Context, taskID primitive.ObjectID)(Task, error)
}

type UserRepository interface{
	Create(c context.Context, user *User) error
	FindByUsername(c context.Context, username string)(User, error)
}

type UserUsecase interface{
	Signup(c context.Context, user *User)error
	Login(c context.Context, request User)(User, string, error)
}