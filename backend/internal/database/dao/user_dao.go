package dao

import (
	"backend/internal/database/models"

	"gorm.io/gorm"
)

type UserDAO struct {
	DB *gorm.DB
}

// GetCount gets the total count of users in the database
func (dao *UserDAO) GetCount() (int64, error) {
	var count int64
	if err := dao.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (dao *UserDAO) Create(user *models.User) error {
	if err := dao.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (dao *UserDAO) GetByID(userID uint) (*models.User, error) {
	var user models.User
	if err := dao.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := dao.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) Update(user *models.User) error {
	if err := dao.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (dao *UserDAO) Delete(userID uint) error {
	if err := dao.DB.Delete(&models.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}
