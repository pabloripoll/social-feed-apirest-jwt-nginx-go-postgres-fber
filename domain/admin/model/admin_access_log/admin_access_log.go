package model

import (
	"time"

	userModel "apirest/domain/user/model"
	"gorm.io/datatypes"
)

// AdminAccessLog maps to the admins_access_logs table.
type AdminAccessLog struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint64         `gorm:"not null;index" json:"user_id"`
	IsTerminated bool           `gorm:"not null;default:false" json:"is_terminated"`
	IsExpired    bool           `gorm:"not null;default:false" json:"is_expired"`
	ExpiresAt    time.Time      `gorm:"not null;index:idx_admins_access_logs_expires_at" json:"expires_at"`
	RefreshCount int            `gorm:"not null;default:0" json:"refresh_count"`
	CreatedAt    time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"not null" json:"updated_at"`
	IPAddress    *string        `gorm:"size:45" json:"ip_address,omitempty"`
	UserAgent    *string        `gorm:"type:text" json:"user_agent,omitempty"`
	RequestsCount int           `gorm:"not null;default:0" json:"requests_count"`
	Payload      datatypes.JSON `gorm:"type:jsonb" json:"payload,omitempty"`
	Token        string         `gorm:"type:text;not null;index:idx_admins_access_logs_token" json:"token"`

	// Associations
	User *userModel.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
}

// TableName sets the exact table name.
func (AdminAccessLog) TableName() string {
	return "admins_access_logs"
}
