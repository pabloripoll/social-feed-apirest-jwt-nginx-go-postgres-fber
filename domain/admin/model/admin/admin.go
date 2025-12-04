package model

import (
	"time"

	userModel "apirest/domain/user/model"
	geoRegionModel "apirest/domain/geo/model/geo_region"
)

// Admin maps to the admins table.
type Admin struct {
	ID                   uint64           `gorm:"primaryKey;autoIncrement" json:"id"`
	UID                  uint64           `gorm:"not null;uniqueIndex:uniq_admins_uid" json:"uid"`
	UserID               uint64           `gorm:"not null;index" json:"user_id"`
	RegionID             uint64           `gorm:"not null;index" json:"region_id"`
	IsActive             bool             `gorm:"not null;default:false" json:"is_active"`
	IsBanned             bool             `gorm:"not null;default:false" json:"is_banned"`
	CreatedAt            time.Time        `gorm:"not null" json:"created_at"`
	UpdatedAt            time.Time        `gorm:"not null" json:"updated_at"`

	// Associations
	User *userModel.User      	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
	Region *geoRegionModel.GeoRegion `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RegionID;references:ID" json:"region,omitempty"`
}

// TableName ensures the struct maps to the exact table name in your SQL.
func (Admin) TableName() string {
	return "admins"
}