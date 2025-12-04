package model

import (
	"time"

	memberNotificationTypeModel "apirest/domain/member/model/member_notification_type"
	userModel "apirest/domain/user/model"
)

// MemberNotification maps to the members_notifications table.
type MemberNotification struct {
	ID                  uint64                                 `gorm:"primaryKey;autoIncrement" json:"id"`
	NotificationTypeID  uint64                                 `gorm:"not null;index" json:"notification_type_id"`
	UserID              uint64                                 `gorm:"not null;index" json:"user_id"`
	IsOpened            bool                                   `gorm:"not null;default:false" json:"is_opened"`
	OpenedAt            time.Time                              `gorm:"not null;index:idx_members_notifications_opened_at" json:"opened_at"`
	CreatedAt           time.Time                              `gorm:"not null;index:idx_members_notifications_created_at" json:"created_at"`
	UpdatedAt           time.Time                              `gorm:"not null" json:"updated_at"`
	Message             string                                 `gorm:"size:512;not null" json:"message"`
	LastMemberUserID    uint64                                 `gorm:"not null;index" json:"last_member_user_id"`
	LastMemberNickname  string                                 `gorm:"size:32;not null" json:"last_member_nickname"`
	LastMemberAvatar    *string                                `gorm:"type:text;default:NULL" json:"last_member_avatar,omitempty"`

	// Associations
	NotificationType *memberNotificationTypeModel.MemberNotificationType	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:NotificationTypeID;references:ID" json:"notification_type,omitempty"`
	User             *userModel.User                                		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID" json:"user,omitempty"`
	LastMemberUser   *userModel.User                                		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:LastMemberUserID;references:ID" json:"last_member_user,omitempty"`
}

// TableName sets the exact table name used by GORM.
func (MemberNotification) TableName() string {
	return "members_notifications"
}
