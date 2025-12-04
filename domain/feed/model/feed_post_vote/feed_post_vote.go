package model

import (
	"time"

	userModel "apirest/domain/user/model"
	feedPostModel "apirest/domain/feed/model/feed_post"
)

// FeedPostVote maps to the feed_posts_votes table.
type FeedPostVote struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint64         `gorm:"not null;index" json:"user_id"`
	PostID       uint64         `gorm:"not null;index" json:"post_id"`
	Up           bool           `gorm:"not null;default:false" json:"up"`
	Down         bool           `gorm:"not null;default:false" json:"down"`
	RefreshCount int            `gorm:"not null;default:0" json:"refresh_count"`
	CreatedAt    time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"not null" json:"updated_at"`

	// Associations
	User 		*userModel.User			`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
	FeedPost    *feedPostModel.FeedPost `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PostID;references:ID" json:"post,omitempty"`
}

// TableName sets the exact table name used by GORM.
func (FeedPostVote) TableName() string {
	return "feed_posts_votes"
}
