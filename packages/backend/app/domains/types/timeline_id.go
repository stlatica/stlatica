//nolint:dupl // TODO: 後ほど共通化する https://github.com/stlatica/stlatica/issues/207
package types

import (
	"database/sql/driver"

	"github.com/friendsofgo/errors"
	"github.com/oklog/ulid/v2"
)

// TimelineID is the type for timeline ID.
type TimelineID ulid.ULID

// NewTimelineID returns a new timeline ID.
func NewTimelineID() TimelineID {
	return TimelineID(ulid.Make())
}

// NewTimelineIDFromString returns a new timeline ID from string.
func NewTimelineIDFromString(timelineIDStr string) (TimelineID, error) {
	timelineID, err := ulid.Parse(timelineIDStr)
	return TimelineID(timelineID), err
}

// String implements fmt.Stringer interface.
func (id TimelineID) String() string {
	return ulid.ULID(id).String()
}

// Value implements driver.Valuer interface.
func (id TimelineID) Value() (driver.Value, error) {
	return id.String(), nil
}

// Scan implements of sql.Scanner interface.
func (id *TimelineID) Scan(src interface{}) error {
	if src == nil {
		return errors.New("TimelineID is not allowed to be nil")
	}

	if u, ok := src.([]byte); ok {
		timelineID, err := ulid.Parse(string(u))
		if err != nil {
			return err
		}
		*id = TimelineID(timelineID)
		return nil
	}
	return errors.New("failed to scan TimelineID")
}
