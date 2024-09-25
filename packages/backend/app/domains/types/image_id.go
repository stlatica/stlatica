//nolint:dupl // TODO: 後ほど共通化する https://github.com/stlatica/stlatica/issues/207
package types

import (
	"database/sql/driver"

	"github.com/friendsofgo/errors"
	"github.com/oklog/ulid/v2"
)

// ImageID is the type for image ID.
type ImageID ulid.ULID

// NewImageID returns a new image ID.
func NewImageID() ImageID {
	return ImageID(ulid.Make())
}

// NewImageIDFromString returns a new image ID from string.
func NewImageIDFromString(imageIDStr string) (ImageID, error) {
	imageID, err := ulid.Parse(imageIDStr)
	return ImageID(imageID), err
}

// String implements fmt.Stringer interface.
func (id ImageID) String() string {
	return ulid.ULID(id).String()
}

// Value implements driver.Valuer interface.
func (id ImageID) Value() (driver.Value, error) {
	return id.String(), nil
}

// Scan implements of sql.Scanner interface.
func (id *ImageID) Scan(src interface{}) error {
	if src == nil {
		return errors.New("ImageID is not allowed to be nil")
	}

	if u, ok := src.([]byte); ok {
		imageID, err := ulid.Parse(string(u))
		if err != nil {
			return err
		}
		*id = ImageID(imageID)
		return nil
	}
	return errors.New("failed to scan ImageID")
}
