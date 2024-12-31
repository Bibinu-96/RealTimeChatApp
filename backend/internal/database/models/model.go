package models

import (
	"strconv"
	"strings"
	"time"
)

type MessageType string

const (
	MessageTypeText MessageType = "text"
	MessageTypeFile MessageType = "file"
)

type User struct {
	UserID          uint      `gorm:"primaryKey"`
	Username        string    `gorm:"size:50;not null;unique"`
	Email           string    `gorm:"size:100;not null;unique"`
	PasswordHash    string    `gorm:"size:255;not null"` // Added PasswordHash field
	PhoneNumber     *string   `gorm:"size:20"`
	Messages        []Message `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
	InteractedUsers string    `gorm:"size:500"` // Comma-separated list of user IDs
}

// Helper methods to manage InteractedUsers as []int64
func (u *User) GetInteractedUsers() ([]int64, error) {
	if u.InteractedUsers == "" {
		return []int64{}, nil
	}
	parts := strings.Split(u.InteractedUsers, ",")
	var result []int64
	for _, p := range parts {
		id, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, id)
	}
	return result, nil
}

func (u *User) SetInteractedUsers(ids []int64) {
	var strIds []string
	for _, id := range ids {
		strIds = append(strIds, strconv.FormatInt(id, 10))
	}
	u.InteractedUsers = strings.Join(strIds, ",")
}

// AddInteractedUser: Adds a user ID to the InteractedUsers field
func (u *User) AddInteractedUser(newID int64) {
	currentIDs, _ := u.GetInteractedUsers()
	// Check if the ID already exists to avoid duplicates
	for _, id := range currentIDs {
		if id == newID {
			return // ID already exists, no need to add
		}
	}
	currentIDs = append(currentIDs, newID) // Add new ID
	u.SetInteractedUsers(currentIDs)       // Update the field
}

// RemoveInteractedUser: Removes a user ID from the InteractedUsers field
func (u *User) RemoveInteractedUser(removeID int64) {
	currentIDs, _ := u.GetInteractedUsers()
	var updatedIDs []int64
	for _, id := range currentIDs {
		if id != removeID {
			updatedIDs = append(updatedIDs, id) // Keep all IDs except the one to remove
		}
	}
	u.SetInteractedUsers(updatedIDs) // Update the field
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
