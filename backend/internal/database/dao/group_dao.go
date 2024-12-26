package dao

import "chatbackend/internal/database/models"

// CreateGroup inserts a new group into the database
func CreateGroup(group *models.Group) error {
	return GetDB().Create(group).Error
}

// FindGroupByID retrieves a group by ID
func FindGroupByID(groupID uint) (*models.Group, error) {
	var group models.Group
	err := GetDB().Preload("Members").First(&group, groupID).Error
	return &group, err
}

// AddGroupMember adds a user to a group
func AddGroupMember(groupID, userID uint) error {
	groupMember := models.GroupMember{
		GroupID: groupID,
		UserID:  userID,
	}
	return GetDB().Create(&groupMember).Error
}

// RemoveGroupMember removes a user from a group
func RemoveGroupMember(groupID, userID uint) error {
	return GetDB().Where("group_id = ? AND user_id = ?", groupID, userID).Delete(&models.GroupMember{}).Error
}
