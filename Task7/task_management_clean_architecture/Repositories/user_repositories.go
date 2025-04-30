package repositories

import (
	"context"
	domain "task_management_clean_architecture/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository{
	return &userRepository{
		database: *db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User)error{
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) FindByUsername(c context.Context, username string)(domain.User, error){
	collection := ur.database.Collection((ur.collection))

	filter := bson.M{"username": username}
	var user domain.User
	
	err := collection.FindOne(c, filter).Decode(&user)
	if err != nil{
		return domain.User{}, err
	}

	return user, nil
}

