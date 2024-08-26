package ports

import "github.com/stlatica/stlatica/packages/backend/app/domains/types"

// FollowsGetParams is the parameters to get follows.
type FollowsGetParams struct {
	UserID           types.UserID
	UserPaginationID types.UserID
	Limit            int
}

// FollowersGetParams is the parameters to get followers.
type FollowersGetParams struct {
	UserID           types.UserID
	UserPaginationID types.UserID
	Limit            int
}
