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
	Password string `json:"-" bson:"password"`
	Role string `json:"role" bson:"role"`
}

type TaskRepository interface{
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string)([]Task, error)
}

type TaskUsecase interface{
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string)([]Task, error)
}

type UserRepository interface{
	Create(c context.Context, user *User) error
	FindByUsername(c context.Context, username string)(User, error)
}

type UserUsecase interface{
	Create(c context.Context, user *User)error
	FindByUsername(c context.Context, username string)(User, error)
}