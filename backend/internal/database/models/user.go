package models

// import "time"

// type User struct {
// 	UserID       uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
// 	Username     string    `gorm:"unique;not null" json:"username"`
// 	Email        string    `gorm:"unique;not null" json:"email"`
// 	PasswordHash string    `gorm:"not null" json:"-"`
// 	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`

// 	// // Relationships
// 	// Groups           []Group       `gorm:"foreignKey:CreatedBy" json:"groups"`             // Groups created by the user
// 	// GroupMembers     []GroupMember `gorm:"foreignKey:UserID" json:"group_members"`         // Groups that the user is a member of
// 	// SentMessages     []Message     `gorm:"foreignKey:SenderID" json:"sent_messages"`       // Messages sent by the user
// 	// ReceivedMessages []Message     `gorm:"foreignKey:ReceiverID" json:"received_messages"` // Messages received by the user
// }
