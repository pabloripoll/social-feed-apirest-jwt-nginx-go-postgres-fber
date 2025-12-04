package model

import (
	"time"

	userModel "apirest/domain/user/model"
)

// MemberFollowing maps to the members_following table.
type MemberFollowing struct {
	ID              uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          uint64         `gorm:"not null;index" json:"user_id"`
	FollowingUserID uint64         `gorm:"not null;index" json:"following_user_id"`
	CreatedAt       time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"not null" json:"updated_at"`

	// Associations to users table (both reference users.id)
	User          *userModel.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
	FollowingUser *userModel.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:FollowingUserID;references:ID" json:"following_user,omitempty"`
}

// TableName sets the exact table name.
func (MemberFollowing) TableName() string {
	return "members_following"
}
