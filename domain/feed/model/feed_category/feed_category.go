package model

import "time"

// FeedCategory maps to the feed_categories table.
type FeedCategory struct {
	ID                          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Key                         string    `gorm:"size:64;not null;uniqueIndex:uniq_feed_categories_key" json:"key"`
	Title                       string    `gorm:"size:64;not null" json:"title"`
	VisitsCount                 int       `gorm:"not null;default:0" json:"visits_count"`
	FeedPostsCount              int       `gorm:"not null;default:0" json:"feed_posts_count"`
	FeedPostsVotesUpCount       int       `gorm:"not null;default:0" json:"feed_posts_votes_up_count"`
	FeedPostsVotesDownCount     int       `gorm:"not null;default:0" json:"feed_posts_votes_down_count"`
	CreatedAt                   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt                   time.Time `gorm:"not null" json:"updated_at"`
}

// TableName returns the exact table name.
func (FeedCategory) TableName() string {
	return "feed_categories"
}
