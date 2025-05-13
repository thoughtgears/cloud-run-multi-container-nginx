package handlers

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/thoughtgears/cloud-run-multi-container-nginx/apis/users/models"
)

var users []models.User

func init() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	userCount := rng.Intn(81) + 20
	users = models.GenerateUser(userCount)
}

func RegisterUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", getUsers)
		userRoutes.GET("/:id", getUserByID)
		userRoutes.POST("", createUser)
		userRoutes.PUT("/:id", updateUser)
		userRoutes.DELETE("/:id", deleteUser)
	}
}

func getUsers(c *gin.Context) {
	c.JSON(200, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			c.JSON(200, user)
			return
		}
	}
	c.JSON(404, gin.H{"message": "User not found"})
}

func createUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = uuid.NewString()
	users = append(users, newUser)
	c.JSON(201, newUser)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			users[i].ID = id
			c.JSON(200, users[i])
			return
		}
	}
	c.JSON(404, gin.H{"message": "User not found"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(200, gin.H{"message": "User deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": "User not found"})
}
