package types

import "github.com/oklog/ulid/v2"

// ImageID is the type for image ID.
type ImageID ulid.ULID

// NewImageID returns a new image ID.
func NewImageID() ImageID {
	return ImageID(ulid.Make())
}

// String implements fmt.Stringer interface.
func (id ImageID) String() string {
	return ulid.ULID(id).String()
}
