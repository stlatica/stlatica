package ports

// FollowsGetParams is the parameters to get follows.
type FollowsGetParams struct {
	PreferredUserID           string
	PreferredUserPaginationID string
	Limit                     int
}

// FollowersGetParams is the parameters to get follows.
type FollowersGetParams struct {
	PreferredUserID           string
	PreferredUserPaginationID string
	Limit                     int
}
