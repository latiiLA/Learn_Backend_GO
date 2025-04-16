package data

import (
	"context"
	"go-auth-task-manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("aetwewereetweqwetwer")

func Register(collection *mongo.Collection, user models.User) (any, error){
	// User registration logic
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func Login(collection *mongo.Collection, user models.User) (any, error) {
	var existingUser models.User
	// User login logic
	err := collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return nil, err
	}

	// Generate JWT
	// Add expiration time (e.g., 24 hours from now)
	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": existingUser.UserID,
		"username": existingUser.Username,
		"role": existingUser.Role,
		"exp": expirationTime,
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}