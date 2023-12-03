//nolint:dupl // TODO: 後ほど共通化する https://github.com/stlatica/stlatica/issues/207
package types

import (
	"database/sql/driver"

	"github.com/friendsofgo/errors"
	"github.com/oklog/ulid/v2"
)

// PlatID is the type for plat ID.
type PlatID ulid.ULID

// NewPlatID returns a new plat ID.
func NewPlatID() PlatID {
	return PlatID(ulid.Make())
}

// NewPlatIDFromString returns a new plat ID from string.
func NewPlatIDFromString(platIDStr string) (PlatID, error) {
	platID, err := ulid.Parse(platIDStr)
	return PlatID(platID), err
}

// String implements fmt.Stringer interface.
func (id PlatID) String() string {
	return ulid.ULID(id).String()
}

// Value implements driver.Valuer interface.
func (id PlatID) Value() (driver.Value, error) {
	return id.String(), nil
}

// Scan implements of sql.Scanner interface.
func (id *PlatID) Scan(src interface{}) error {
	if src == nil {
		return errors.New("PlatID is not allowed to be nil")
	}

	if u, ok := src.([]byte); ok {
		platID, err := ulid.Parse(string(u))
		if err != nil {
			return err
		}
		*id = PlatID(platID)
		return nil
	}
	return errors.New("failed to scan PlatID")
}
