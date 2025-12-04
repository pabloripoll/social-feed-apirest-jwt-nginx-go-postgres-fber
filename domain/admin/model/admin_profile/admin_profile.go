package model

import (
	"time"

	userModel "apirest/domain/user/model"
)

// AdminProfile maps to the admins_profile table.
type AdminProfile struct {
	ID        uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64          `gorm:"not null;index" json:"user_id"`
	Nickname  string          `gorm:"size:32;not null;uniqueIndex:uniq_admins_profile_nickname" json:"nickname"`
	Avatar    *string         `gorm:"type:text;default:NULL" json:"avatar,omitempty"`
	CreatedAt time.Time       `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"not null" json:"updated_at"`

	// Association to User. Cascade on delete to match the SQL FK.
	User *userModel.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
}

// TableName sets the exact table name.
func (AdminProfile) TableName() string {
	return "admins_profile"
}