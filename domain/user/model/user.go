package model

import "time"

// User maps to the users table.
type User struct {
	ID              uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Role            string     `gorm:"size:16;not null" json:"role"`
	Email           string     `gorm:"size:64;not null;uniqueIndex:uniq_users_email;index:idx_users_email" json:"email"`
	EmailVerifiedAt *time.Time `gorm:"default:NULL" json:"email_verified_at,omitempty"`
	Password        string     `gorm:"column:password;size:128;not null" json:"-"`
	CreatedAt       time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"not null" json:"updated_at"`
}

// TableName ensures the struct maps to the exact table name in your SQL.
func (User) TableName() string {
	return "users"
}
