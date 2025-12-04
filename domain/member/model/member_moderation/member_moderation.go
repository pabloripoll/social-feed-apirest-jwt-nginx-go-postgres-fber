package model

import (
	"time"

	userModel "apirest/domain/user/model"
	memberModerationTypeModel "apirest/domain/member/model/member_moderation_type"
	feedPostModel "apirest/domain/feed/model/feed_post"
)

// MemberModeration maps to the members_moderations table.
type MemberModeration struct {
	ID               uint64                      `gorm:"primaryKey;autoIncrement" json:"id"`
	AdminUserID      uint64                      `gorm:"not null;index" json:"admin_user_id"`
	TypeID           uint64                      `gorm:"not null;index" json:"type_id"`
	IsApplied        bool                        `gorm:"not null;default:false" json:"is_applied"`
	ExpiresAt        time.Time                   `gorm:"not null;index:idx_members_moderations_expires_at" json:"expires_at"`
	IsOnMember       bool                        `gorm:"not null;default:false" json:"is_on_member"`
	IsOnPost         bool                        `gorm:"not null;default:false" json:"is_on_post"`
	MemberUserID     *uint64                     `gorm:"index" json:"member_user_id,omitempty"`
	MemberFeedPostID *uint64                     `gorm:"index" json:"member_feed_post_id,omitempty"`
	CreatedAt        time.Time                   `gorm:"not null" json:"created_at"`
	UpdatedAt        time.Time                   `gorm:"not null" json:"updated_at"`

	// Associations
	AdminUser      *userModel.User                    				`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AdminUserID;references:ID" json:"admin_user,omitempty"`
	MemberUser     *userModel.User                    				`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:MemberUserID;references:ID" json:"member_user,omitempty"`
	ModerationType *memberModerationTypeModel.MemberModerationType 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TypeID;references:ID" json:"type,omitempty"`
	MemberFeedPost *feedPostModel.FeedPost            				`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:MemberFeedPostID;references:ID" json:"member_feed_post,omitempty"`
}

// TableName sets the exact table name used by GORM.
func (MemberModeration) TableName() string {
	return "members_moderations"
}
