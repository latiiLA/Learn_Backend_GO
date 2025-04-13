package main

import (
	"context"
	"fmt"
	"log"
	"task_manager/db"
	"task_manager/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Defer disconnect (will execute when main() exits)
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("Warning: failed to disconnect from MongoDB: %v", err)
		} else {
			log.Println("MongoDB connection closed")
		}
	}()

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	
	collections := &db.DBCollections{
		Tasks: client.Database("task_manager").Collection("tasks"),
	}

	r := router.RouterSetup(collections)
	
	if err := r.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	
}