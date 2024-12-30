package dao

import (
	"backend/internal/database/models"

	"gorm.io/gorm"
)

type MessageDAO struct {
	DB *gorm.DB
}

func (dao *MessageDAO) Create(message *models.Message) error {
	if err := dao.DB.Create(message).Error; err != nil {
		return err
	}
	return nil
}

func (dao *MessageDAO) GetByID(messageID uint) (*models.Message, error) {
	var message models.Message
	if err := dao.DB.First(&message, messageID).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (dao *MessageDAO) GetMessagesBySender(senderID uint, limit, offset int) ([]models.Message, error) {
	var messages []models.Message
	if err := dao.DB.Where("sender_id = ?", senderID).Limit(limit).Offset(offset).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (dao *MessageDAO) GetMessagesForGroup(groupID uint, limit, offset int) ([]models.Message, error) {
	var messages []models.Message
	if err := dao.DB.Where("group_id = ?", groupID).Limit(limit).Offset(offset).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (dao *MessageDAO) GetMessagesForReceiver(receiverID uint, limit, offset int) ([]models.Message, error) {
	var messages []models.Message
	if err := dao.DB.Where("receiver_id = ?", receiverID).Limit(limit).Offset(offset).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (dao *MessageDAO) Update(message *models.Message) error {
	if err := dao.DB.Save(message).Error; err != nil {
		return err
	}
	return nil
}

func (dao *MessageDAO) Delete(messageID uint) error {
	if err := dao.DB.Delete(&models.Message{}, messageID).Error; err != nil {
		return err
	}
	return nil
}

// GetMessagesInGroup retrieves paginated messages in a specific group
func (dao *MessageDAO) GetMessagesInGroup(groupID uint, page, pageSize int) ([]models.Message, error) {
	var messages []models.Message
	offset := (page - 1) * pageSize // Calculate the offset for the page

	// Fetch paginated messages for the group
	if err := dao.DB.Preload("Sender").Preload("Receiver").
		Where("group_id = ?", groupID).
		Limit(pageSize).Offset(offset).
		Order("created_at desc").
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

// GetMessagesBetweenSenderAndReceiver retrieves paginated messages between a specific sender and receiver
func (dao *MessageDAO) GetMessagesBetweenSenderAndReceiver(senderID, receiverID uint, page, pageSize int) ([]models.Message, error) {
	var messages []models.Message
	offset := (page - 1) * pageSize // Calculate the offset for the page

	// Fetch paginated messages between the sender and receiver
	if err := dao.DB.Preload("Sender").Preload("Receiver").
		Where("sender_id = ? AND receiver_id = ?", senderID, receiverID).
		Limit(pageSize).Offset(offset).
		Order("created_at desc").
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

// GetCount gets the total count of messages in the database
func (dao *MessageDAO) GetCount() (int64, error) {
	var count int64
	if err := dao.DB.Model(&models.Message{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// MessageDAO
func (dao *MessageDAO) GetCountBetweenSenderAndReceiver(senderID, receiverID uint) (int64, error) {
	var count int64
	if err := dao.DB.Model(&models.Message{}).
		Where("sender_id = ? AND receiver_id = ?", senderID, receiverID).
		Count(&count).
		Error; err != nil {
		return 0, err
	}
	return count, nil
}

// MessageDAO
func (dao *MessageDAO) GetCountForGroup(groupID uint) (int64, error) {
	var count int64
	if err := dao.DB.Model(&models.Message{}).
		Where("group_id = ?", groupID).
		Count(&count).
		Error; err != nil {
		return 0, err
	}
	return count, nil
}
