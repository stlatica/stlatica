// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entities

import (
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// UserBase is UserBase entity object.
type UserBase struct { // ulid
	UserID types.UserID `json:"user_id"`
	// preferred user id
	PreferredUserID string `json:"preferred_user_id"`
	// preferred user name
	PreferredUserName string `json:"preferred_user_name"`
	// Unix time
	RegisteredAt types.UnixTime `json:"registered_at"`
	// user is public
	IsPublic bool `json:"is_public"`
	// mail address
	MailAddress string `json:"mail_address"`
	// Unix time
	CreatedAt types.UnixTime `json:"created_at"`
	// Unix time
	UpdatedAt types.UnixTime `json:"updated_at"`

	R *UserBaseR `json:"-"`
}

// UserBaseR is where relationships are stored.
type UserBaseR struct {
	Plats PlatBaseSlice `json:"Plats"`
}

// UserBaseSlice is an alias for a slice of pointers to UserBase.
// This should almost always be used instead of []UserBase.
type UserBaseSlice []*UserBase

// GetUserID is get user_id value, if receiver is nil, returns the specified value.
func (m *UserBase) GetUserID() types.UserID {
	if m == nil {
		return types.UserID{}
	}
	return m.UserID
}

// GetPreferredUserID is get preferred_user_id value, if receiver is nil, returns the specified value.
func (m *UserBase) GetPreferredUserID() string {
	if m == nil {
		return ""
	}
	return m.PreferredUserID
}

// GetPreferredUserName is get preferred_user_name value, if receiver is nil, returns the specified value.
func (m *UserBase) GetPreferredUserName() string {
	if m == nil {
		return ""
	}
	return m.PreferredUserName
}

// GetRegisteredAt is get registered_at value, if receiver is nil, returns the specified value.
func (m *UserBase) GetRegisteredAt() types.UnixTime {
	if m == nil {
		return types.UnixTime(0)
	}
	return m.RegisteredAt
}

// GetIsPublic is get is_public value, if receiver is nil, returns the specified value.
func (m *UserBase) GetIsPublic() bool {
	if m == nil {
		return false
	}
	return m.IsPublic
}

// GetMailAddress is get mail_address value, if receiver is nil, returns the specified value.
func (m *UserBase) GetMailAddress() string {
	if m == nil {
		return ""
	}
	return m.MailAddress
}

// GetCreatedAt is get created_at value, if receiver is nil, returns the specified value.
func (m *UserBase) GetCreatedAt() types.UnixTime {
	if m == nil {
		return types.UnixTime(0)
	}
	return m.CreatedAt
}

// GetUpdatedAt is get updated_at value, if receiver is nil, returns the specified value.
func (m *UserBase) GetUpdatedAt() types.UnixTime {
	if m == nil {
		return types.UnixTime(0)
	}
	return m.UpdatedAt
}

// Plats is get Plats relationship, if receiver is nil, returns empty PlatBaseSlice.
func (o *UserBase) Plats() PlatBaseSlice {
	if o == nil {
		return PlatBaseSlice{}
	}
	return o.R.Plats
}
