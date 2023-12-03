package types

import (
	"database/sql/driver"

	"github.com/friendsofgo/errors"
	"github.com/oklog/ulid/v2"
)

// UserID is the type for user ID.
type UserID ulid.ULID

// NewUserID returns a new user ID.
func NewUserID() UserID {
	return UserID(ulid.Make())
}

// NewUserIDFromString returns a new user ID from string.
func NewUserIDFromString(userIDStr string) (UserID, error) {
	userID, err := ulid.Parse(userIDStr)
	return UserID(userID), err
}

// String implements fmt.Stringer interface.
func (id UserID) String() string {
	return ulid.ULID(id).String()
}

// Value implements driver.Valuer interface.
func (id UserID) Value() (driver.Value, error) {
	return id.String(), nil
}

// Scan implements of sql.Scanner interface.
func (id *UserID) Scan(src interface{}) error {
	if src == nil {
		return errors.New("UserID is not allowed to be nil")
	}

	if u, ok := src.([]byte); ok {
		userID, err := ulid.Parse(string(u))
		if err != nil {
			return err
		}
		*id = UserID(userID)
		return nil
	}
	return errors.New("failed to scan UserID")
}
