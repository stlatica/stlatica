package ports

// FollowsGetParams is the parameters to get follows.
type FollowsGetParams struct {
	PreferredUserID           string
	PreferredUserPaginationID string
	Limit                     uint64
}
