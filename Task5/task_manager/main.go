package main

import (
	"context"
	"fmt"
	"log"
	"task_manager/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.TODO()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Warning: failed to disconnect from MongoDB: %v", err)
		} else {
			log.Println("MongoDB connection closed")
		}
	}()

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Ping error:", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Directly create and pass the controller
	taskCollection := client.Database("task_manager").Collection("tasks")

	r := router.RouterSetup(taskCollection)

	if err := r.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
