//nolint:dupl // TODO: 後ほど共通化する https://github.com/stlatica/stlatica/issues/207
package types

import (
	"database/sql/driver"

	"github.com/friendsofgo/errors"
	"github.com/oklog/ulid/v2"
)

// ActorID is the type for actor ID.
type ActorID ulid.ULID

// NewActorID returns a new actor ID.
func NewActorID() ActorID {
	return ActorID(ulid.Make())
}

// NewActorIDFromString returns a new actor ID from string.
func NewActorIDFromString(actorIDStr string) (ActorID, error) {
	actorID, err := ulid.Parse(actorIDStr)
	return ActorID(actorID), err
}

// String implements fmt.Stringer interface.
func (id ActorID) String() string {
	return ulid.ULID(id).String()
}

// Value implements driver.Valuer interface.
func (id ActorID) Value() (driver.Value, error) {
	return id.String(), nil
}

// Scan implements of sql.Scanner interface.
func (id *ActorID) Scan(src interface{}) error {
	if src == nil {
		return errors.New("ActorID is not allowed to be nil")
	}

	if u, ok := src.([]byte); ok {
		actorID, err := ulid.Parse(string(u))
		if err != nil {
			return err
		}
		*id = ActorID(actorID)
		return nil
	}
	return errors.New("failed to scan ActorID")
}
