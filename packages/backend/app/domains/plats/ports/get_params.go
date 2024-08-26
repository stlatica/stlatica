package ports

import "github.com/stlatica/stlatica/packages/backend/app/domains/types"

// PlatsGetParams is the parameters to get plats.
type PlatsGetParams struct {
	UserIDs  []*types.UserID
	ToDate   types.UnixTime
	FromDate types.UnixTime
	Limit    int
}
