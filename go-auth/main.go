package main

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]*User)
var jwtSecret = []byte("aetwewereetweqwetwer")

func AuthMiddleware() gin.HandlerFunc{
	return func (c *gin.Context)  {
		// implementation of JWT validation logic
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer"{
			c.JSON(401, gin.H{"error" : "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid{
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.GET("/", AuthMiddleware(), func (c *gin.Context)  {
		c.JSON(200, gin.H{"message": "Welcome to the GO Authentication and Authorization tutorial!"})
	})

	router.POST("/register",  func (c *gin.Context)  {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil{
			c.JSON(400, gin.H{"error": "Invalid request payload"})
		}

		// User registration logic
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil{
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		user.Password = string(hashedPassword)
		users[user.Email] = &user

		c.JSON(200, gin.H{"message": "User registered successfully"})

	})

	router.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil{
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		// User login logic
		existingUser, ok := users[user.Email]
		if !ok || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil{
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		// Generate JWT

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": existingUser.ID,
			"email": existingUser.Email,
		})

		jwtToken, err := token.SignedString(jwtSecret)
		if err != nil{
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(200, gin.H{"message": "User logged in successfull", "token": jwtToken})
	})

	router.Run()
}