package model

import (
	"time"

	userModel "apirest/domain/user/model"
	feedPostModel "apirest/domain/feed/model/feed_post"
)

// FeedPostVisit maps to the feed_posts_visits table.
type FeedPostVisit struct {
	ID             uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         uint64         `gorm:"not null;index" json:"user_id"`
	PostID         uint64         `gorm:"not null;index" json:"post_id"`
	VisitorUserID  uint64         `gorm:"not null;index" json:"visitor_user_id"`
	CreatedAt      time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"not null" json:"updated_at"`

	// Associations (ON DELETE CASCADE)
	User        *userModel.User  		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
	FeedPost    *feedPostModel.FeedPost `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PostID;references:ID" json:"post,omitempty"`
	VisitorUser *userModel.User  		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:VisitorUserID;references:ID" json:"visitor_user,omitempty"`
}

// TableName returns the exact table name used by the SQL DDL.
func (FeedPostVisit) TableName() string {
	return "feed_posts_visits"
}
