// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entities

import (
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// PlatBase is PlatBase entity object.
type PlatBase struct { // ulid
	PlatID types.PlatID `json:"plat_id"`
	// user_id
	UserID types.UserID `json:"user_id"`
	// body text
	Content string `json:"content"`
	// Unix time
	CreatedAt types.UnixTime `json:"created_at"`
	// Unix time
	UpdatedAt types.UnixTime `json:"updated_at"`

	R *PlatBaseR `json:"-"`
}

// PlatBaseR is where relationships are stored.
type PlatBaseR struct {
	User *UserBase `json:"User"`
}

// PlatBaseSlice is an alias for a slice of pointers to PlatBase.
// This should almost always be used instead of []PlatBase.
type PlatBaseSlice []*PlatBase

// GetPlatID is get plat_id value, if receiver is nil, returns the specified value.
func (m *PlatBase) GetPlatID() types.PlatID {
	if m == nil {
		return types.PlatID{}
	}
	return m.PlatID
}

// GetUserID is get user_id value, if receiver is nil, returns the specified value.
func (m *PlatBase) GetUserID() types.UserID {
	if m == nil {
		return types.UserID{}
	}
	return m.UserID
}

// GetContent is get content value, if receiver is nil, returns the specified value.
func (m *PlatBase) GetContent() string {
	if m == nil {
		return ""
	}
	return m.Content
}

// GetCreatedAt is get created_at value, if receiver is nil, returns the specified value.
func (m *PlatBase) GetCreatedAt() types.UnixTime {
	if m == nil {
		return types.UnixTime(0)
	}
	return m.CreatedAt
}

// GetUpdatedAt is get updated_at value, if receiver is nil, returns the specified value.
func (m *PlatBase) GetUpdatedAt() types.UnixTime {
	if m == nil {
		return types.UnixTime(0)
	}
	return m.UpdatedAt
}
