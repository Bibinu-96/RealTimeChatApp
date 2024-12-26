package dao

import "chatbackend/internal/database/models"

// SendMessage creates a new message (either direct or group)
func SendMessage(message *models.Message) error {
	return GetDB().Create(message).Error
}

// GetMessagesForUser retrieves messages sent to a specific user (both direct and group)
func GetMessagesForUser(userID uint) ([]models.Message, error) {
	var messages []models.Message
	err := GetDB().
		Preload("Sender").
		Preload("Receiver").
		Preload("Group").
		Where("receiver_id = ? OR group_id IN (SELECT group_id FROM group_members WHERE user_id = ?)", userID, userID).
		Find(&messages).Error
	return messages, err
}

// GetMessagesForGroup retrieves all messages in a specific group
func GetMessagesForGroup(groupID uint) ([]models.Message, error) {
	var messages []models.Message
	err := GetDB().
		Preload("Sender").
		Where("group_id = ?", groupID).
		Find(&messages).Error
	return messages, err
}
