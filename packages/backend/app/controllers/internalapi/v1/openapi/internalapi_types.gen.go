// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package openapi

import (
	"time"
)

const (
	BearerScopes = "Bearer.Scopes"
)

// Defines values for ErrorResponseCode.
const (
	BADREQUEST          ErrorResponseCode = "BAD_REQUEST"
	CONFLICT            ErrorResponseCode = "CONFLICT"
	INTERNALSERVERERROR ErrorResponseCode = "INTERNAL_SERVER_ERROR"
	MISSINGPARAMETER    ErrorResponseCode = "MISSING_PARAMETER"
	NOTFOUND            ErrorResponseCode = "NOT_FOUND"
	SERVICEUNAVAILABLE  ErrorResponseCode = "SERVICE_UNAVAILABLE"
	UNAUTHORIZED        ErrorResponseCode = "UNAUTHORIZED"
	UNPROCESSABLEENTITY ErrorResponseCode = "UNPROCESSABLE_ENTITY"
)

// Defines values for TimelineType.
const (
	TimelineTypeFollowing TimelineType = "following"
	TimelineTypeHome      TimelineType = "home"
	TimelineTypeLocal     TimelineType = "local"
)

// Defines values for GetTimelineByQueryParamsType.
const (
	GetTimelineByQueryParamsTypeFollowing GetTimelineByQueryParamsType = "following"
	GetTimelineByQueryParamsTypeHome      GetTimelineByQueryParamsType = "home"
	GetTimelineByQueryParamsTypeLocal     GetTimelineByQueryParamsType = "local"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Code    ErrorResponseCode `json:"code"`
	Message string            `json:"message"`
}

// ErrorResponseCode defines model for ErrorResponse.Code.
type ErrorResponseCode string

// Plat plat
type Plat struct {
	// Content platの本文
	Content string `json:"content"`

	// CreatedAt platが作成された日時(ISO8601)
	CreatedAt time.Time `json:"created_at"`

	// FavoriteCount platがお気に入りされた数
	FavoriteCount *int `json:"favorite_count,omitempty"`

	// ImageUrls platに添付された画像のURL
	ImageUrls *[]string `json:"image_urls,omitempty"`

	// PlatId platを識別するための一意のID
	PlatId PlatID `json:"plat_id"`

	// ReplatCount platがリプラットされた数
	ReplatCount *int `json:"replat_count,omitempty"`

	// ReplyCount platに対するリプライの数
	ReplyCount *int `json:"reply_count,omitempty"`

	// UserId userを識別するための一意のID
	UserId *UserID `json:"user_id,omitempty"`
}

// PlatID platを識別するための一意のID
type PlatID = string

// PlatPost plat
type PlatPost struct {
	// Content platの本文
	Content string `json:"content"`

	// UserId userを識別するための一意のID
	UserId UserID `json:"user_id"`
}

// User user
type User struct {
	// FollowersCount follower数
	FollowersCount int `json:"followers_count"`

	// FollowingCount following数
	FollowingCount int `json:"following_count"`

	// Icon ユーザのアイコン画像のURL
	Icon string `json:"icon"`

	// IsPublic 公開アカウントであるかどうか \
	// external apiのmanuallyApprovesFollowersと同一の値となる
	IsPublic bool `json:"is_public"`

	// Summary ユーザのプロフィール
	Summary string `json:"summary"`

	// UserId userを識別するための一意のID
	UserId UserID `json:"user_id"`

	// Username 画面上に表示されるユーザ名
	Username string `json:"username"`
}

// UserID userを識別するための一意のID
type UserID = string

// PlatId defines model for plat_id.
type PlatId = string

// TimelineFromDate この日時以降のplatを取得する
// from_dateおよびto_dateは以下の制約を持つ
// - from_dateもしくはto_dateいずれかの指定は必須
// - from_dateとto_dateを同時に指定することはできない
// - to_dateと同時に指定された場合、status code 422を返す
type TimelineFromDate = time.Time

// TimelineId defines model for timeline_id.
type TimelineId = string

// TimelineLimit 取得するplatの最大数
// デフォルトは100
type TimelineLimit = int

// TimelineToDate この日時以前のplatを取得する
// from_dateおよびto_dateは以下の制約を持つ
// - from_dateもしくはto_dateいずれかの指定は必須
// - from_dateとto_dateを同時に指定することはできない
// - from_dateと同時に指定された場合、status code 422を返す
type TimelineToDate = time.Time

// TimelineType timelineの種類
// - home: 指定したuser_idのplatの配列
// - following: 指定したuser_idのfollowingのplatの配列
// - local: インスタンス内の全てのplatの配列
type TimelineType string

// TimelineUserId defines model for timeline_user_id.
type TimelineUserId = string

// UserId defines model for user_id.
type UserId = string

// UploadImageTextBody defines parameters for UploadImage.
type UploadImageTextBody = string

// LoginJSONBody defines parameters for Login.
type LoginJSONBody struct {
	Password        *string `json:"password,omitempty"`
	PreferredUserId *string `json:"preferred_user_id,omitempty"`
}

// GetTimelineByQueryParams defines parameters for GetTimelineByQuery.
type GetTimelineByQueryParams struct {
	UserId   TimelineUserId               `form:"user_id" json:"user_id"`
	Type     GetTimelineByQueryParamsType `form:"type" json:"type"`
	Limit    *TimelineLimit               `form:"limit,omitempty" json:"limit,omitempty"`
	FromDate *TimelineFromDate            `form:"from_date,omitempty" json:"from_date,omitempty"`
	ToDate   *TimelineToDate              `form:"to_date,omitempty" json:"to_date,omitempty"`
}

// GetTimelineByQueryParamsType defines parameters for GetTimelineByQuery.
type GetTimelineByQueryParamsType string

// GetUsersParams defines parameters for GetUsers.
type GetUsersParams struct {
	// UserName ユーザ名
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// UploadImageTextRequestBody defines body for UploadImage for text/plain ContentType.
type UploadImageTextRequestBody = UploadImageTextBody

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody LoginJSONBody

// PostPlatJSONRequestBody defines body for PostPlat for application/json ContentType.
type PostPlatJSONRequestBody = PlatPost

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody
