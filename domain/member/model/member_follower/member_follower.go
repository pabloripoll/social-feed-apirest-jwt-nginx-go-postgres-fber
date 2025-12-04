package model

import (
	"time"

	userModel "apirest/domain/user/model"
)

// MemberFollower maps to the members_followers table.
type MemberFollower struct {
	ID             uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         uint64         `gorm:"not null;index" json:"user_id"`
	FollowerUserID uint64         `gorm:"not null;index" json:"follower_user_id"`
	CreatedAt      time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"not null" json:"updated_at"`

	// Associations to users table (both reference users.id)
	User         *userModel.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
	FollowerUser *userModel.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:FollowerUserID;references:ID" json:"follower_user,omitempty"`
}

// TableName sets the exact table name.
func (MemberFollower) TableName() string {
	return "members_followers"
}
