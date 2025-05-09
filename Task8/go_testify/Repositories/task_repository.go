package repositories

import (
	"context"
	domain "task_management_clean_architecture/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	database mongo.Database
	collection string
}

func NewTaskRepository(db *mongo.Database, collection string) domain.TaskRepository{
	return &taskRepository{
		database: *db,
		collection: collection,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task)error{
	collection := tr.database.Collection(tr.collection)

	_, err := collection.InsertOne(c, task)

	return err
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID primitive.ObjectID)([]domain.Task, error){
	collection := tr.database.Collection(tr.collection)

	var tasks []domain.Task

	cursor, err := collection.Find(c, bson.M{"userID": userID})
	if err != nil{
		return nil, err
	}

	err = cursor.All(c, &tasks)
	if tasks == nil{
		return []domain.Task{}, err
	}

	return tasks, err
}

func (tr *taskRepository) FetchAll(c context.Context)([]domain.Task, error){
	collection := tr.database.Collection(tr.collection)

	var tasks []domain.Task

	cursor, err := collection.Find(c, bson.M{})
	if err != nil{
		return nil, err
	}
	defer cursor.Close(c)
	
	if err := cursor.All(c, &tasks); err != nil{
		return nil, err
	}

	return tasks, err
}

func (tr *taskRepository) UpdateByTaskID(c context.Context, taskID primitive.ObjectID, updatedData *domain.Task)(int64, int64, error){
	collection := tr.database.Collection(tr.collection)

	result, err := collection.UpdateOne(
		c,
		bson.M{"_id": taskID},
		bson.M{"$set": updatedData},
	)

	if err != nil{
		return 0, 0, err
	}
	return result.MatchedCount, result.ModifiedCount, nil
}

func (tr *taskRepository) DeleteByTaskID(c context.Context, taskID primitive.ObjectID)(int64, error){
	collection := tr.database.Collection(tr.collection)

	result, err := collection.DeleteOne(
		c, 
		bson.M{"_id": taskID},
	)

	if err != nil{
		return 0, err
	}

	return result.DeletedCount, nil
}

func (tr *taskRepository) FetchByTaskID(c context.Context, taskID primitive.ObjectID)(domain.Task, error){
	collection := tr.database.Collection(tr.collection)

	var task domain.Task

	err := collection.FindOne(c, bson.M{"_id": taskID}).Decode(&task)

	if err != nil{
		return domain.Task{}, err
	}

	return task, nil
}