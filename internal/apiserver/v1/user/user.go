package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// User represents body of User request and response.
type User struct {
	// User's name.
	// Required: true
	Name string `json:"name"`

	// User's nickname.
	// Required: true
	Nickname string `json:"nickname"`

	// User's address.
	Address string `json:"address"`

	// User's email.
	Email string `json:"email"`
}

var users []*User

// Create a user in memory.
func Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 10001})
		return
	}

	for _, u := range users {
		if u.Name == user.Name {
			c.JSON(http.StatusBadRequest,
				gin.H{"message": fmt.Sprintf("user %s already exist", user.Name), "code": 10001})
			return
		}
	}

	users = append(users, &user)
	c.JSON(http.StatusOK, user)
}

// Get return the detail information for a user.
func Get(c *gin.Context) {
	username := c.Param("name")
	for _, u := range users {
		if u.Name == username {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s not exist", username), "code": 10002})
}
