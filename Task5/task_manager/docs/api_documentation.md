How to connect to mongodb from Go project.

This documentation contains step by step instruction on how to connect mongodb from go project

1. Install Mongodb
   The first step of all is installing mongo db on your personal computer or using online existing mongo db services like MongoAtlas.

   Mongo Atlas has free access with limited functionlity. So you can start with that if you do not want to install mongodb for now.
   Go to this link and create an account.
   https://www.mongodb.com/cloud/atlas/register

   If you opted for installing mongodb on your personal computer download mongo db from the following link and follow the installation procedure.
   https://www.mongodb.com/try/download/community

2. Install Go lang
   I think you have already did install go lang if you are doing this task. But incase you did not install it follow this documentation.
   https://go.dev/doc/

3. Start Go Project

   - Create Project name as a folder
     mkdir your project_name
   - Run the following command to initiate go package
     go mod init "example.com/package_name"
   - After that create main.go
     touch main.go

4. Install gin package

   - To install gin package use the following command
     go get github.com/gin-gonic/gin

5. Install necessary packages

   - Install mongo-driver package
     go get go.mongodb.org/mongo-driver
   - Install mongo-bson package
     go.mongodb.org/mongo-driver/bson \
     go.mongodb.org/mongo-driver/mongo/options

6. Connect to mongo DB

   - To connect to mongodb you have to write the following code in main.go. You can also do it in other files but it is easy to setup in main.go. If you are working on big projects you can opt to do it in other files.

   import necessary packages
   for example: - import ("go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options")
   Then follow the following steps:

   1.1 create a context
   for example: - ctx := context.TODO()
   1.2 create a client options
   for example: - clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
   If you have secured your database you may need to add that here, as you did on the ApplyURI.
   1.3 Connect to your mongo db
   for example: - client, err := mongo.Connect(ctx, clientOptions)
   1.4 Check if you are connected or not
   for example: -
   if err != nil {
   log.Fatal("Mongo connection error:", err)
   }
   1.5 If you reached here that means you are connected to mongodb. If you did not stumble to an error.
   1.6 Now you can create a database and database table(collection) to your mongodb
   for example: - taskCollection := client.Database("task_manager").Collection("tasks")
