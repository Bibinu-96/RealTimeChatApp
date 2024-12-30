package router

import (
	"net/http"

	"backend/internal/businesslogic"

	"github.com/gin-gonic/gin"
)

func SetupGinRouter() *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Create API routes
	api := r.Group("/api")
	{
		api.POST("/register", RegisterUser)
		api.POST("/login", LoginUser)
		api.GET("/users/connected", GetConnectedUsers)
		api.GET("/groups/use", GetGroups)
		api.GET("/messages/direct", GetDirectMessages)
		api.GET("/messages/group", GetGroupMessages)
	}
	return r
}

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {

	var register businesslogic.RegisterUser
	err := c.BindJSON(&register)
	if err == nil {
		userService := businesslogic.GetUserServiceInstance()
		err = userService.RegisterUserForApp(register)
		if err != nil {
			c.JSON(500, gin.H{"status": "internal server error"})
		} else {
			c.JSON(201, gin.H{"status": register.USERNAME})
		}
		return
	}

}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	// Implement your logic here
	var loginUser businesslogic.LOGIN
	err := c.BindJSON(&loginUser)
	if err == nil {
		userService := businesslogic.GetUserServiceInstance()
		err = userService.LoginUserForApp(loginUser)
		if err != nil {
			c.JSON(500, gin.H{"status": "internal server error"})
		} else {
			c.JSON(200, gin.H{"status": loginUser.EMAIL})
		}
	}

}

// GetConnectedUsers retrieves connected users
func GetConnectedUsers(c *gin.Context) {
	// Implement your logic here
	c.JSON(http.StatusOK, gin.H{"message": "Get connected users logic not implemented from gin"})
}

// GetGroups retrieves groups
func GetGroups(c *gin.Context) {
	// Implement your logic here
	c.JSON(http.StatusOK, gin.H{"message": "Get groups logic not implemented from gin"})
}

// GetDirectMessages retrieves direct messages
func GetDirectMessages(c *gin.Context) {
	// Implement your logic here
	c.JSON(http.StatusOK, gin.H{"message": "Get direct messages logic not implemented from gin"})
}

// GetGroupMessages retrieves group messages
func GetGroupMessages(c *gin.Context) {
	// Implement your logic here
	c.JSON(http.StatusOK, gin.H{"message": "Get group messages logic not implemented from gin"})
}
