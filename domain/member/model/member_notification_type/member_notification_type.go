package model

import "time"

// MemberNotificationType maps to the members_notification_types table.
type MemberNotificationType struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Key             string    `gorm:"size:64;not null;uniqueIndex:uniq_members_notification_types_key" json:"key"`
	Title           string    `gorm:"size:64;not null" json:"title"`
	MessageSingular string    `gorm:"size:512;not null" json:"message_singular"`
	MessageMultiple string    `gorm:"size:512;not null" json:"message_multiple"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"not null" json:"updated_at"`
}

// TableName sets the exact table name used by GORM.
func (MemberNotificationType) TableName() string {
	return "members_notification_types"
}
