package model

import (
	"time"

	userModel "apirest/domain/user/model"
	geoRegionModel "apirest/domain/geo/model/geo_region"
)

// Member maps to the members table.
type Member struct {
	ID                   uint64           `gorm:"primaryKey;autoIncrement" json:"id"`
	UID                  uint64           `gorm:"not null;uniqueIndex:uniq_members_uid" json:"uid"`
	UserID               uint64           `gorm:"not null;index" json:"user_id"`
	RegionID             uint64           `gorm:"not null;index" json:"region_id"`
	IsActive             bool             `gorm:"not null;default:false" json:"is_active"`
	IsBanned             bool             `gorm:"not null;default:false" json:"is_banned"`
	FollowingCount       int              `gorm:"not null;default:0" json:"following_count"`
	FollowersCount       int              `gorm:"not null;default:0" json:"followers_count"`
	PostsCount           int              `gorm:"not null;default:0" json:"posts_count"`
	PostsVotesUpCount    int              `gorm:"not null;default:0" json:"posts_votes_up_count"`
	PostsVotesDownCount  int              `gorm:"not null;default:0" json:"posts_votes_down_count"`
	CreatedAt            time.Time        `gorm:"not null" json:"created_at"`
	UpdatedAt            time.Time        `gorm:"not null" json:"updated_at"`

	// Associations
	User *userModel.User      	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
	Region *geoRegionModel.GeoRegion `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RegionID;references:ID" json:"region,omitempty"`
}

// TableName ensures the struct maps to the exact table name in your SQL.
func (Member) TableName() string {
	return "members"
}