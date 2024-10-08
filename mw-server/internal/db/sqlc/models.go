// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type PricingPlanType string

const (
	PricingPlanTypeFree    PricingPlanType = "free"
	PricingPlanTypeStarter PricingPlanType = "starter"
	PricingPlanTypePro     PricingPlanType = "pro"
)

func (e *PricingPlanType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PricingPlanType(s)
	case string:
		*e = PricingPlanType(s)
	default:
		return fmt.Errorf("unsupported scan type for PricingPlanType: %T", src)
	}
	return nil
}

type NullPricingPlanType struct {
	PricingPlanType PricingPlanType `json:"pricing_plan_type"`
	Valid           bool            `json:"valid"` // Valid is true if PricingPlanType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPricingPlanType) Scan(value interface{}) error {
	if value == nil {
		ns.PricingPlanType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PricingPlanType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPricingPlanType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PricingPlanType), nil
}

type Comment struct {
	Uuid          pgtype.UUID      `json:"uuid"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	Description   string           `json:"description"`
	OwnerUuid     pgtype.UUID      `json:"owner_uuid"`
	DayReportUuid pgtype.UUID      `json:"day_report_uuid"`
}

type CompositeWay struct {
	ChildUuid  pgtype.UUID `json:"child_uuid"`
	ParentUuid pgtype.UUID `json:"parent_uuid"`
}

type DayReport struct {
	Uuid      pgtype.UUID      `json:"uuid"`
	WayUuid   pgtype.UUID      `json:"way_uuid"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type FavoriteUser struct {
	DonorUserUuid    pgtype.UUID `json:"donor_user_uuid"`
	AcceptorUserUuid pgtype.UUID `json:"acceptor_user_uuid"`
}

type FavoriteUsersWay struct {
	UserUuid pgtype.UUID `json:"user_uuid"`
	WayUuid  pgtype.UUID `json:"way_uuid"`
}

type FormerMentorsWay struct {
	FormerMentorUuid pgtype.UUID `json:"former_mentor_uuid"`
	WayUuid          pgtype.UUID `json:"way_uuid"`
}

type FromUserMentoringRequest struct {
	UserUuid pgtype.UUID `json:"user_uuid"`
	WayUuid  pgtype.UUID `json:"way_uuid"`
}

type JobDone struct {
	Uuid          pgtype.UUID      `json:"uuid"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	Description   string           `json:"description"`
	Time          int32            `json:"time"`
	OwnerUuid     pgtype.UUID      `json:"owner_uuid"`
	DayReportUuid pgtype.UUID      `json:"day_report_uuid"`
}

type JobDonesJobTag struct {
	JobDoneUuid pgtype.UUID `json:"job_done_uuid"`
	JobTagUuid  pgtype.UUID `json:"job_tag_uuid"`
}

type JobTag struct {
	Uuid        pgtype.UUID `json:"uuid"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Color       string      `json:"color"`
	WayUuid     pgtype.UUID `json:"way_uuid"`
}

type MentorUsersWay struct {
	UserUuid pgtype.UUID `json:"user_uuid"`
	WayUuid  pgtype.UUID `json:"way_uuid"`
}

type Metric struct {
	Uuid             pgtype.UUID      `json:"uuid"`
	CreatedAt        pgtype.Timestamp `json:"created_at"`
	UpdatedAt        pgtype.Timestamp `json:"updated_at"`
	Description      string           `json:"description"`
	IsDone           bool             `json:"is_done"`
	DoneDate         pgtype.Timestamp `json:"done_date"`
	MetricEstimation int32            `json:"metric_estimation"`
	WayUuid          pgtype.UUID      `json:"way_uuid"`
}

type Plan struct {
	Uuid          pgtype.UUID      `json:"uuid"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	Description   string           `json:"description"`
	Time          int32            `json:"time"`
	OwnerUuid     pgtype.UUID      `json:"owner_uuid"`
	IsDone        bool             `json:"is_done"`
	DayReportUuid pgtype.UUID      `json:"day_report_uuid"`
}

type PlansJobTag struct {
	PlanUuid   pgtype.UUID `json:"plan_uuid"`
	JobTagUuid pgtype.UUID `json:"job_tag_uuid"`
}

type Problem struct {
	Uuid          pgtype.UUID      `json:"uuid"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	Description   string           `json:"description"`
	IsDone        bool             `json:"is_done"`
	OwnerUuid     pgtype.UUID      `json:"owner_uuid"`
	DayReportUuid pgtype.UUID      `json:"day_report_uuid"`
}

type ProfileSetting struct {
	Uuid           pgtype.UUID      `json:"uuid"`
	PricingPlan    PricingPlanType  `json:"pricing_plan"`
	ExpirationDate pgtype.Timestamp `json:"expiration_date"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
	OwnerUuid      pgtype.UUID      `json:"owner_uuid"`
}

type ToUserMentoringRequest struct {
	UserUuid pgtype.UUID `json:"user_uuid"`
	WayUuid  pgtype.UUID `json:"way_uuid"`
}

type User struct {
	Uuid        pgtype.UUID      `json:"uuid"`
	Name        string           `json:"name"`
	Email       string           `json:"email"`
	Description string           `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	ImageUrl    string           `json:"image_url"`
	IsMentor    bool             `json:"is_mentor"`
}

type UserTag struct {
	Uuid pgtype.UUID `json:"uuid"`
	Name string      `json:"name"`
}

type UsersUserTag struct {
	UserUuid    pgtype.UUID `json:"user_uuid"`
	UserTagUuid pgtype.UUID `json:"user_tag_uuid"`
}

type Way struct {
	Uuid              pgtype.UUID      `json:"uuid"`
	Name              string           `json:"name"`
	GoalDescription   string           `json:"goal_description"`
	UpdatedAt         pgtype.Timestamp `json:"updated_at"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	EstimationTime    int32            `json:"estimation_time"`
	OwnerUuid         pgtype.UUID      `json:"owner_uuid"`
	CopiedFromWayUuid pgtype.UUID      `json:"copied_from_way_uuid"`
	IsCompleted       bool             `json:"is_completed"`
	IsPrivate         bool             `json:"is_private"`
}

type WayCollection struct {
	Uuid      pgtype.UUID      `json:"uuid"`
	OwnerUuid pgtype.UUID      `json:"owner_uuid"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	Name      string           `json:"name"`
	Type      string           `json:"type"`
}

type WayCollectionsWay struct {
	WayCollectionUuid pgtype.UUID `json:"way_collection_uuid"`
	WayUuid           pgtype.UUID `json:"way_uuid"`
}

type WayTag struct {
	Uuid pgtype.UUID `json:"uuid"`
	Name string      `json:"name"`
}

type WaysWayTag struct {
	WayUuid    pgtype.UUID `json:"way_uuid"`
	WayTagUuid pgtype.UUID `json:"way_tag_uuid"`
}
