package model

import "time"

// FeedReportType maps to the feed_report_types table.
type FeedReportType struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Key         string     `gorm:"size:64;not null" json:"key"`
	Title       string     `gorm:"size:64;not null" json:"title"`
	Description *string    `gorm:"size:256" json:"description,omitempty"`
	Level       int16      `gorm:"not null;default:0" json:"level"`
	Position    int16      `gorm:"not null;default:0" json:"position"`
	CreatedAt   time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"not null" json:"updated_at"`
}

// TableName returns the exact table name used by GORM.
func (FeedReportType) TableName() string {
	return "feed_report_types"
}
