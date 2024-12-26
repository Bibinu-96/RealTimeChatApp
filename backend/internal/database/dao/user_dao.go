package dao

import "chatbackend/internal/database/models"

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	return GetDB().Create(user).Error
}

// FindUserByID retrieves a user by ID
func FindUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := GetDB().First(&user, userID).Error
	return &user, err
}

// FindUserByEmail retrieves a user by email
func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := GetDB().Where("email = ?", email).First(&user).Error
	return &user, err
}

// DeleteUser removes a user by ID
func DeleteUser(userID uint) error {
	return GetDB().Delete(&models.User{}, userID).Error
}
