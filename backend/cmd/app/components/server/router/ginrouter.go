package router

import (
	"net/http"

	"backend/internal/businesslogic"
	"backend/internal/database/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupGinRouter() *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// CORS Middleware to allow all methods and all origins
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"}, // Allow all origins
		AllowMethods:  []string{"*"}, // Allow all methods
		AllowHeaders:  []string{"*"}, // Allow all headers
		ExposeHeaders: []string{"*"}, // Expose all headers
	}))
	// Create API routes
	api := r.Group("/api")
	{
		api.POST("/register", RegisterUser)
		api.POST("/login", LoginUser)
		api.POST("/users/interaction", CheckAuth, AddInteractedUser)
		api.DELETE("/users/interaction", CheckAuth, DeleteInteractedUsers)
		api.POST("/users/interactions", CheckAuth, GetInteractedUsers) // get of all interacted users
		api.GET("/groups/", CheckAuth, GetGroups)
		api.GET("/messages/direct", CheckAuth, GetDirectMessages)
		api.GET("/messages/group", CheckAuth, GetGroupMessages)
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
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(201, gin.H{"status": "success"})
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
		token, err := userService.LoginUserForApp(loginUser)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"token": token})
		}
	}

}

// GetInteractedUsers retrieves interacted users for the given user
func GetInteractedUsers(c *gin.Context) {
	var currentUser *models.User
	user, exists := c.Get("currentUser")
	if exists {
		currentUser = user.(*models.User)
	}
	var paginationInfo businesslogic.PaginationInfo
	err := c.BindJSON(&paginationInfo)

	if err == nil {
		chatSevice := businesslogic.GetChatServiceInstance()
		interactedUsers, total, err := chatSevice.GetInteractedUsers(currentUser, paginationInfo)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK,
				gin.H{"users": interactedUsers, "total": total})
		}

	}
}

// GetInteractedUsers retrieves interacted users for the given user
func AddInteractedUser(c *gin.Context) {
	var currentUser *models.User
	user, exists := c.Get("currentUser")
	if exists {
		currentUser = user.(*models.User)
	}

	var interactedUser businesslogic.InteractedUser
	err := c.BindJSON(&interactedUser)

	if err == nil {
		chatSevice := businesslogic.GetChatServiceInstance()
		err = chatSevice.AddUserToInteractedListOfCurrentUser(currentUser, interactedUser)
		c.JSON(400, gin.H{"error": err.Error()})

	}
	c.JSON(201, gin.H{"status": "success"})
}

// GetInteractedUsers retrieves interacted users for the given user
func DeleteInteractedUsers(c *gin.Context) {
	var currentUser *models.User
	user, exists := c.Get("currentUser")
	if exists {
		currentUser = user.(*models.User)
	}
	var interactedUser businesslogic.InteractedUser
	err := c.BindJSON(&interactedUser)

	if err == nil {
		chatSevice := businesslogic.GetChatServiceInstance()
		err = chatSevice.RemoveUserFromInteractedListOfCurrentUser(currentUser, interactedUser)
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"status": "success"})
}

// GetGroups retrieves groups
func GetGroups(c *gin.Context) {

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
