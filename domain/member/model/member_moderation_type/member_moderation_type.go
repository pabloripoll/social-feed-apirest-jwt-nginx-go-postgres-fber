package model

import "time"

// MemberModerationType maps to the members_moderation_types table.
type MemberModerationType struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Key         string    `gorm:"size:64;not null;uniqueIndex:uniq_members_moderation_types_key" json:"key"`
	Title       string    `gorm:"size:64;not null" json:"title"`
	Description string    `gorm:"size:256;not null" json:"description"`
	Position    int16     `gorm:"not null;default:0" json:"position"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}

// TableName sets the exact table name used by GORM.
func (MemberModerationType) TableName() string {
	return "members_moderation_types"
}
