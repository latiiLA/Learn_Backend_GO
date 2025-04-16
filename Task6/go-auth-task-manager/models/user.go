package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string       `json:"username"`
	Password string       `json:"-"`
	Role string	`json:"role"`
}