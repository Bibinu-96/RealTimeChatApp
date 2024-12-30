package models

// import "time"

// type Message struct {
// 	MessageID   uint      `gorm:"primaryKey;autoIncrement" json:"message_id"`
// 	SenderID    uint      `gorm:"not null" json:"sender_id"` // Foreign key to User (sender)
// 	ReceiverID  *uint     `json:"receiver_id"`               // Foreign key to User (for direct messages, nullable)
// 	GroupID     *uint     `json:"group_id"`                  // Foreign key to Group (for group messages, nullable)
// 	Content     string    `gorm:"type:text;not null" json:"content"`
// 	MessageType string    `gorm:"type:varchar(10);not null" json:"message_type"` // 'direct' or 'group'
// 	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`

// 	// Relationships
// 	Sender   User   `gorm:"foreignKey:SenderID" json:"sender"`
// 	Receiver *User  `gorm:"foreignKey:ReceiverID" json:"receiver,omitempty"`
// 	Group    *Group `gorm:"foreignKey:GroupID" json:"group,omitempty"`
// }
