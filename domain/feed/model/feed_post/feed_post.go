package model

import (
	"time"

	userModel "apirest/domain/user/model"
	geoRegionModel "apirest/domain/geo/model/geo_region"
	feedCategoryModel "apirest/domain/feed/model/feed_category"
)

// FeedPost maps to the feed_posts table.
type FeedPost struct {
	ID               uint64                  `gorm:"primaryKey;autoIncrement" json:"id"`
	UID              uint64                  `gorm:"not null;uniqueIndex:uniq_posts_uid" json:"uid"`
	UserID           uint64                  `gorm:"not null;index" json:"user_id"`
	RegionID         uint64                  `gorm:"not null;index" json:"region_id"`
	CategoryID       uint64                  `gorm:"not null;index" json:"category_id"`
	IsActive         bool                    `gorm:"not null;default:false" json:"is_active"`
	IsDraft          bool                    `gorm:"not null;default:false" json:"is_draft"`
	IsBanned         bool                    `gorm:"not null;default:false" json:"is_banned"`
	VisitsCount      int                     `gorm:"not null;default:0" json:"visits_count"`
	ReportsCount     int                     `gorm:"not null;default:0" json:"reports_count"`
	VotesUpCount     int                     `gorm:"not null;default:0" json:"votes_up_count"`
	VotesDownCount   int                     `gorm:"not null;default:0" json:"votes_down_count"`
	Title            *string                 `gorm:"size:128" json:"title,omitempty"`
	Slug             *string                 `gorm:"size:128" json:"slug,omitempty"`
	Summary          *string                 `gorm:"size:256" json:"summary,omitempty"`
	Article          *string                 `gorm:"type:text" json:"article,omitempty"`
	CreatedAt        time.Time               `gorm:"not null" json:"created_at"`
	UpdatedAt        time.Time               `gorm:"not null" json:"updated_at"`

	// Associations
	User     *userModel.User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
	Region   *geoRegionModel.GeoRegion `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RegionID;references:ID" json:"region,omitempty"`
	Category *feedCategoryModel.FeedCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CategoryID;references:ID" json:"category,omitempty"`
}

// TableName sets the exact table name used by GORM.
func (FeedPost) TableName() string {
	return "feed_posts"
}
