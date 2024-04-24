// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entities

import (
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// User is User entity object.
type User struct { // ulid
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
}

// GetUserID is get user_id value, if receiver is nil, returns the specified value.
func (m *User) GetUserID() types.UserID {
	if m == nil {
		return types.UserID{}
	}
	return m.UserID
}

// GetPreferredUserID is get preferred_user_id value, if receiver is nil, returns the specified value.
func (m *User) GetPreferredUserID() string {
	if m == nil {
		return ""
	}
	return m.PreferredUserID
}

// GetPreferredUserName is get preferred_user_name value, if receiver is nil, returns the specified value.
func (m *User) GetPreferredUserName() string {
	if m == nil {
		return ""
	}
	return m.PreferredUserName
}

// GetRegisteredAt is get registered_at value, if receiver is nil, returns the specified value.
func (m *User) GetRegisteredAt() types.UnixTime {
	if m == nil {
		return types.UnixTime(0)
	}
	return m.RegisteredAt
}

// GetIsPublic is get is_public value, if receiver is nil, returns the specified value.
func (m *User) GetIsPublic() bool {
	if m == nil {
		return false
	}
	return m.IsPublic
}

// GetMailAddress is get mail_address value, if receiver is nil, returns the specified value.
func (m *User) GetMailAddress() string {
	if m == nil {
		return ""
	}
	return m.MailAddress
}

// GetCreatedAt is get created_at value, if receiver is nil, returns the specified value.
func (m *User) GetCreatedAt() types.UnixTime {
	if m == nil {
		return types.UnixTime(0)
	}
	return m.CreatedAt
}

// GetUpdatedAt is get updated_at value, if receiver is nil, returns the specified value.
func (m *User) GetUpdatedAt() types.UnixTime {
	if m == nil {
		return types.UnixTime(0)
	}
	return m.UpdatedAt
}