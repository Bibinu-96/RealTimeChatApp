package models

import "time"

type GroupMember struct {
	GroupMemberID uint      `gorm:"primaryKey;autoIncrement" json:"group_member_id"`
	GroupID       uint      `gorm:"not null" json:"group_id"` // Foreign key to Group
	UserID        uint      `gorm:"not null" json:"user_id"`  // Foreign key to User
	JoinedAt      time.Time `gorm:"autoCreateTime" json:"joined_at"`

	// Relationships
	Group Group `gorm:"foreignKey:GroupID" json:"group"`
	User  User  `gorm:"foreignKey:UserID" json:"user"`
}
