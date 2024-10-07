package context_test

import (
	"context"
	"reflect"
	"testing"

	domaincontexts "github.com/stlatica/stlatica/packages/backend/app/domains/contexts"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// TestUserIDValue tests the UserIDValue function
func TestUserIDValue(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name         string
		args         args
		wantID       string
		wantIsExists bool
	}{
		{
			name: "normal",
			args: args{
				ctx:    context.Background(),
				userID: "01J7WX1N1KMWCZAAJVK0J14J0G",
			},
			wantID:       "01J7WX1N1KMWCZAAJVK0J14J0G",
			wantIsExists: true,
		},
		{
			name: "context is nil",
			args: args{
				ctx:    nil,
				userID: "01J7WX1N1KMWCZAAJVK0J14J0G",
			},
			wantID:       "00000000000000000000000000",
			wantIsExists: false,
		},
		{
			name: "user_id is not set",
			args: args{
				ctx:    context.Background(),
				userID: "",
			},
			wantID:       "",
			wantIsExists: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argsUserID, _ := types.NewUserIDFromString(tt.args.userID)
			wantUserID, _ := types.NewUserIDFromString(tt.wantID)

			var ctx context.Context
			if tt.args.ctx != nil && tt.args.userID != "" {
				ctx = domaincontexts.SetUserID(tt.args.ctx, argsUserID)
			}

			gotID, gotIsExists := domaincontexts.UserIDValue(ctx)
			if !reflect.DeepEqual(gotID, wantUserID) {
				t.Errorf("UserIDValue() gotID = %v, want %v", gotID, tt.wantID)
			}
			if gotIsExists != tt.wantIsExists {
				t.Errorf("UserIDValue() gotIsExists = %v, want %v", gotIsExists, tt.wantIsExists)
			}
		})
	}
}

// TestDeleteUserID tests the DeleteUserID function
func TestDeleteUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name         string
		args         args
		wantValue    string
		wantIsExists bool
	}{
		{
			name: "normal",
			args: args{
				ctx:    context.Background(),
				userID: "01J7WX1N1KMWCZAAJVK0J14J0G",
			},
			wantValue:    "00000000000000000000000000",
			wantIsExists: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argsUserID, _ := types.NewUserIDFromString(tt.args.userID)
			ctx := domaincontexts.SetUserID(tt.args.ctx, argsUserID)
			ctx = domaincontexts.DeleteUserID(ctx)
			gotID, gotIsExists := domaincontexts.UserIDValue(ctx)
			if gotID.String() != tt.wantValue {
				t.Errorf("DeleteUserID() gotID = %v, want %v", gotID, tt.wantValue)
			}
			if gotIsExists != tt.wantIsExists {
				t.Errorf("DeleteUserID() gotIsExists = %v, want %v", gotIsExists, tt.wantIsExists)
			}
		})
	}
}
