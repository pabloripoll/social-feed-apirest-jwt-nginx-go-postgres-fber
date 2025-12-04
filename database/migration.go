package database

import (
	geoContinentModel "apirest/domain/geo/model/geo_continent"
	geoRegionModel "apirest/domain/geo/model/geo_region"
	userModel "apirest/domain/user/model"
	memberModel "apirest/domain/member/model/member"
	memberProfileModel "apirest/domain/member/model/member_profile"
	memberAccessLogModel "apirest/domain/member/model/member_access_log"
	memberFollowingModel "apirest/domain/member/model/member_following"
	memberFollowerModel "apirest/domain/member/model/member_follower"
	memberActivationCodeModel "apirest/domain/member/model/member_activation_code"
	feedCategoryModel "apirest/domain/feed/model/feed_category"
	feedPostModel "apirest/domain/feed/model/feed_post"
	adminModel "apirest/domain/admin/model/admin"
	adminProfileModel "apirest/domain/admin/model/admin_profile"
	adminAccessLogModel "apirest/domain/admin/model/admin_access_log"
	memberNotificationTypeModel "apirest/domain/member/model/member_notification_type"
	memberNotificationModel "apirest/domain/member/model/member_notification"
	memberModerationTypeModel "apirest/domain/member/model/member_moderation_type"
	memberModerationModel "apirest/domain/member/model/member_moderation"
	feedReportTypeModel "apirest/domain/feed/model/feed_report_type"
	feedReportModel "apirest/domain/feed/model/feed_report"
	feedPostVoteModel "apirest/domain/feed/model/feed_post_vote"
	feedPostVisitModel "apirest/domain/feed/model/feed_post_visit"

	"fmt"
	"gorm.io/gorm"
)

// Schema migration using the GORM
// The order matters due to foreign key dependencies
func RunMigrations(db *gorm.DB) error {
	models := []interface{}{
		&geoContinentModel.GeoContinent{},
		&geoRegionModel.GeoRegion{},
		&userModel.User{},
		&memberModel.Member{},
		&memberProfileModel.MemberProfile{},
		&memberAccessLogModel.MemberAccessLog{},
		&memberFollowingModel.MemberFollowing{},
		&memberFollowerModel.MemberFollower{},
		&memberActivationCodeModel.MemberActivationCode{},
		&feedCategoryModel.FeedCategory{},
		&feedPostModel.FeedPost{},
		&adminModel.Admin{},
		&adminProfileModel.AdminProfile{},
		&adminAccessLogModel.AdminAccessLog{},
		&memberNotificationTypeModel.MemberNotificationType{},
		&memberNotificationModel.MemberNotification{},
		&memberModerationTypeModel.MemberModerationType{},
		&memberModerationModel.MemberModeration{},
		&feedReportTypeModel.FeedReportType{},
		&feedReportModel.FeedReport{},
		&feedPostVoteModel.FeedPostVote{},
		&feedPostVisitModel.FeedPostVisit{},
	}

	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			// DB/GORM error (e.g. PostgreSQL constraint/index error)
			return fmt.Errorf("auto-migrate failed for %T: %w", m, err)
		}
	}

	fmt.Println("Database migrated.")
	return nil
}
