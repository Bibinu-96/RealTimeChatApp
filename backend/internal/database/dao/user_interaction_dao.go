package dao

import (
	"time"

	"backend/internal/database/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var userInteractionInstance *UserInteractionDAO

type UserInteractionDAO struct {
	db *gorm.DB
}

func GetUserInteractionDAO() *UserInteractionDAO {
	once.Do(func() {
		userInteractionInstance = &UserInteractionDAO{
			db: GetDB(),
		}
	})
	return userInteractionInstance
}

// Insert an interaction
func (dao *UserInteractionDAO) InsertInteraction(userID, interactedUserID uint) error {

	interaction := models.UserInteraction{
		UserID:           userID,
		InteractedUserID: interactedUserID,
		LastInteraction:  time.Now(),
	}

	return dao.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "interacted_user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"last_interaction"}),
	}).Create(&interaction).Error
}

// Check if an interaction exists
func (dao *UserInteractionDAO) InteractionExists(userID, interactedUserID uint) (bool, error) {

	var count int64
	err := dao.db.Model(&models.UserInteraction{}).
		Where("user_id = ? AND interacted_user_id = ?", userID, interactedUserID).
		Count(&count).Error

	return count > 0, err
}

// Delete an interaction
func (dao *UserInteractionDAO) DeleteInteraction(userID, interactedUserID uint) error {

	return dao.db.Where("user_id = ? AND interacted_user_id = ?", userID, interactedUserID).
		Delete(&models.UserInteraction{}).Error
}

// Get all interacted users (paginated)
func (dao *UserInteractionDAO) GetInteractedUsers(userID uint, page, pageSize int) ([]models.UserInteraction, int64, error) {
	var interactions []models.UserInteraction
	var total int64

	err := dao.db.Model(&models.UserInteraction{}).
		Where("user_id = ? OR interacted_user_id = ?", userID, userID).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = dao.db.Where("user_id = ? OR interacted_user_id = ?", userID, userID).
		Offset(offset).Limit(pageSize).
		Find(&interactions).Error

	return interactions, total, err
}
