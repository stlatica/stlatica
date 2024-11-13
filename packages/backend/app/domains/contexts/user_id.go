package contexts

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

type contextKey string

const userIDKey contextKey = "user_id"

// SetUserID sets the user ID in the context.
func SetUserID(ctx context.Context, userID types.UserID) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// DeleteUserID deletes the user ID from the context.
func DeleteUserID(ctx context.Context) context.Context {
	return context.WithValue(ctx, userIDKey, types.UserID{})
}

// UserIDValue retrieves the user ID from the context.
func UserIDValue(ctx context.Context) (userID types.UserID, isExists bool) {
	if ctx == nil {
		return types.UserID{}, false
	}

	ctxUserID, isExist := ctx.Value(userIDKey).(types.UserID)
	if isExist {
		return ctxUserID, true
	}

	return types.UserID{}, false
}
