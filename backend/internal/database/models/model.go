package models

import (
	"time"
)

type MessageType string

const (
	MessageTypeText MessageType = "text"
	MessageTypeFile MessageType = "file"
)

type User struct {
	UserID       uint      `gorm:"primaryKey"`
	Username     string    `gorm:"size:50;not null;unique"`
	Email        string    `gorm:"size:100;not null;unique"`
	PasswordHash string    `gorm:"size:255;not null"` // Added PasswordHash field
	PhoneNumber  *string   `gorm:"size:20"`
	Messages     []Message `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
}

type Group struct {
	GroupID   uint   `gorm:"primaryKey"`
	GroupName string `gorm:"size:50;not null"`

	Messages []Message `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
	Members  []User    `gorm:"many2many:group_members;"`
}

type Message struct {
	MessageID   uint   `gorm:"primaryKey"`
	SenderID    uint   `gorm:"not null"`
	ReceiverID  *uint  `gorm:"index"`
	GroupID     *uint  `gorm:"index"`
	MessageType string `gorm:"size:10;not null;default:'text'"` // Replaced `ENUM` with `string`
	Content     *string
	FileType    string
	FilePath    string
	FileName    string
	Delivered   bool      `gorm:"default:false"` // Added Delivered field
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	Sender      User      `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
	Receiver    *User     `gorm:"foreignKey:ReceiverID;constraint:OnDelete:SET NULL"`
}

// Custom validation for messages
