package dao

import (
	"backend/internal/database/models"

	"gorm.io/gorm"
)

type GroupDAO struct {
	DB *gorm.DB
}

func (dao *GroupDAO) Create(group *models.Group) error {
	if err := dao.DB.Create(group).Error; err != nil {
		return err
	}
	return nil
}

// GetCount gets the total count of groups in the database
func (dao *GroupDAO) GetCount() (int64, error) {
	var count int64
	if err := dao.DB.Model(&models.Group{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (dao *GroupDAO) GetByID(groupID uint) (*models.Group, error) {
	var group models.Group
	if err := dao.DB.Preload("Members").Preload("Messages").First(&group, groupID).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (dao *GroupDAO) GetByName(groupName string) (*models.Group, error) {
	var group models.Group
	if err := dao.DB.Where("group_name = ?", groupName).First(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (dao *GroupDAO) AddMember(groupID uint, userID uint) error {
	var group models.Group
	if err := dao.DB.First(&group, groupID).Error; err != nil {
		return err
	}

	var user models.User
	if err := dao.DB.First(&user, userID).Error; err != nil {
		return err
	}

	// Add user to group members
	if err := dao.DB.Model(&group).Association("Members").Append(&user); err != nil {
		return err
	}

	return nil
}

func (dao *GroupDAO) RemoveMember(groupID uint, userID uint) error {
	var group models.Group
	if err := dao.DB.First(&group, groupID).Error; err != nil {
		return err
	}

	var user models.User
	if err := dao.DB.First(&user, userID).Error; err != nil {
		return err
	}

	// Remove user from group members
	if err := dao.DB.Model(&group).Association("Members").Delete(&user); err != nil {
		return err
	}

	return nil
}

func (dao *GroupDAO) Delete(groupID uint) error {
	if err := dao.DB.Delete(&models.Group{}, groupID).Error; err != nil {
		return err
	}
	return nil
}
