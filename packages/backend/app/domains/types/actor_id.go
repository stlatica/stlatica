package types

import "github.com/oklog/ulid/v2"

// ActorID is the type for actor ID.
type ActorID ulid.ULID

// NewActorID returns a new actor ID.
func NewActorID() ActorID {
	return ActorID(ulid.Make())
}
