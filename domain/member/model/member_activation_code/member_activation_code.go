package model

import (
	"time"

	userModel "apirest/domain/user/model"
)

// MemberActivationCode maps to the members_activation_codes table.
type MemberActivationCode struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string         `gorm:"size:64;not null;uniqueIndex:uniq_members_activation_codes_code" json:"code"`
	UserID    uint64         `gorm:"not null;index" json:"user_id"`
	IsActive  bool           `gorm:"not null;default:false" json:"is_active"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null" json:"updated_at"`

	// Association to users table. Cascade on delete to match the SQL FK.
	User *userModel.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
}

// TableName sets the exact table name.
func (MemberActivationCode) TableName() string {
	return "members_activation_codes"
}
