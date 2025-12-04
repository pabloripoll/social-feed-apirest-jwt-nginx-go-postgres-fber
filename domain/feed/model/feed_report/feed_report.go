package model

import (
	"time"

	userModel "apirest/domain/user/model"
	memberModerationModel "apirest/domain/member/model/member_moderation"
	feedPostModel "apirest/domain/feed/model/feed_post"
	feedReportTypeModel "apirest/domain/feed/model/feed_report_type"
)

// FeedReport maps to the feed_reports table.
type FeedReport struct {
	ID                uint64                        `gorm:"primaryKey;autoIncrement" json:"id"`
	TypeID            uint64                        `gorm:"not null;index" json:"type_id"`
	ReporterUserID    uint64                        `gorm:"not null;index" json:"reporter_user_id"`
	ReporterMessage   *string                       `gorm:"size:256" json:"reporter_message,omitempty"`
	InReview          bool                          `gorm:"not null;default:false" json:"in_review"`
	InReviewSince     time.Time                     `gorm:"not null" json:"in_review_since"`
	IsClosed          bool                          `gorm:"not null;default:false" json:"is_closed"`
	ClosedAt          time.Time                     `gorm:"not null" json:"closed_at"`
	ModerationID      uint64                        `gorm:"not null;index" json:"moderation_id"`
	MemberUserID      uint64                        `gorm:"not null;index" json:"member_user_id"`
	MemberFeedPostID  uint64                        `gorm:"not null;index" json:"member_feed_post_id"`
	CreatedAt         time.Time                     `gorm:"not null;index:idx_feed_reports_created_at" json:"created_at"`
	UpdatedAt         time.Time                     `gorm:"not null" json:"updated_at"`

	// Associations
	FeedReportType *feedReportTypeModel.FeedReportType		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TypeID;references:ID" json:"type,omitempty"`
	ReporterUser   *userModel.User                        	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ReporterUserID;references:ID" json:"reporter_user,omitempty"`
	Moderation     *memberModerationModel.MemberModeration 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ModerationID;references:ID" json:"moderation,omitempty"`
	MemberUser     *userModel.User                        	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:MemberUserID;references:ID" json:"member_user,omitempty"`
	MemberFeedPost *feedPostModel.FeedPost                	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:MemberFeedPostID;references:ID" json:"member_feed_post,omitempty"`
}

// TableName returns the exact table name used by the SQL DDL.
func (FeedReport) TableName() string {
	return "feed_reports"
}
