package types

import "github.com/oklog/ulid/v2"

type ActorID ulid.ULID

func NewActorID() ActorID {
	return ActorID(ulid.Make())
}
