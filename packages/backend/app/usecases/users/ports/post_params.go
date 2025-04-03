package ports

import "github.com/stlatica/stlatica/packages/backend/app/domains/types"

// FollowPostParams is the parameters to post follow.
type FollowPostParams struct {
	UserID       types.UserID
	FollowUserID types.UserID
}
