package main

import (
	"context"
	"fmt"
	"log"
	"task_management_clean_architecture/Delivery/routers"
	"time"

	"github.com/gin-gonic/gin"
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
	// taskCollection := client.Database("task_manager").Collection("tasks")
	// userCollection := client.Database("task_manager").Collection("users")
	db := client.Database("task_manager_db")
	jwtSecret := "Thisismysecretcode"
	timeout := 2 * time.Second
	
	r := gin.Default()
	routers.RouterSetup(jwtSecret, timeout, db, r)

	if err := r.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
