package types

import (
	"database/sql/driver"
	"fmt"
)

// UnixTime is the type for unix time.
type UnixTime uint64

// Value implements driver.Valuer interface.
func (ut UnixTime) Value() (driver.Value, error) {
	return int64(ut), nil
}

// Scan implements of sql.Scanner interface.
func (ut *UnixTime) Scan(src interface{}) error {
	if src == nil {
		*ut = 0
		return nil
	}

	if u, ok := src.(int64); ok {
		*ut = UnixTime(u)
		return nil
	}

	return fmt.Errorf("failed to scan UnixTime: %v", src)
}
