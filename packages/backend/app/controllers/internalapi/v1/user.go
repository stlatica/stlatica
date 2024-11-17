package v1

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/users"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/users/ports"
)

type userController struct {
	userUseCase users.UserUseCase
}

// GetUserResponse is the response of GetUser.
type GetUserResponse struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	RegisteredAt string `json:"registered_at"`
	IsPublic     bool   `json:"is_public"`
	MailAddress  string `json:"mail_address"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// CreateUserResponse is the response of CreateUser.
type CreateUserResponse struct {
	UserID string `json:"user_id"`
}

// GetFollowResponse is the response of GetFollows.
type GetFollowResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Summary  string `json:"summary"`
	Icon     string `json:"icon"`
	IsPublic bool   `json:"is_public"`
}

// GetFollowerResponse is the response of GetFollowers.
type GetFollowerResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Summary  string `json:"summary"`
	Icon     string `json:"icon"`
	IsPublic bool   `json:"is_public"`
}

// GetUser converts request data and calls usecase to get actor.
func (c *userController) GetUser(ectx echo.Context, userID string) (*GetUserResponse, error) {
	user, err := c.userUseCase.GetUserByPreferredUserID(ectx.Request().Context(), userID)
	if err != nil {
		return &GetUserResponse{}, err
	}

	return &GetUserResponse{
		UserID:       user.GetPreferredUserID(),
		Username:     user.GetPreferredUserName(),
		RegisteredAt: user.GetRegisteredAt().ConvertToTime().String(),
		IsPublic:     user.GetIsPublic(),
		MailAddress:  user.GetMailAddress(),
		CreatedAt:    user.GetCreatedAt().ConvertToTime().String(),
		UpdatedAt:    user.GetUpdatedAt().ConvertToTime().String(),
	}, nil
}

// CreateUser creates a new user.
func (c *userController) CreateUser(ectx echo.Context, userName string, preferredUserID string, mailAddress string,
	iconImageIDStr string) (*CreateUserResponse, error) {
	iconImageID, err := types.NewImageIDFromString(iconImageIDStr)
	if err != nil {
		return nil, err
	}
	user, err := c.userUseCase.CreateUser(
		ectx.Request().Context(), userName, preferredUserID, mailAddress, iconImageID)
	if err != nil {
		return nil, err
	}
	return &CreateUserResponse{
		UserID: user.GetPreferredUserID(),
	}, nil
}

func (c *userController) GetFollows(ectx echo.Context,
	userIDStr string, userPaginationID *string, limit *int) ([]*GetFollowResponse, error) {
	var limitValue int
	var userPaginationIDStr string
	if limit != nil {
		limitValue = *limit
	}
	if userPaginationID == nil {
		userPaginationIDStr = ""
	} else {
		userPaginationIDStr = *userPaginationID
	}

	getParams := ports.FollowsGetParams{
		PreferredUserID:           userIDStr,
		PreferredUserPaginationID: userPaginationIDStr,
		Limit:                     limitValue,
	}
	follows, err := c.userUseCase.GetFollows(ectx.Request().Context(), getParams)
	if err != nil {
		return nil, err
	}

	followResponses := make([]*GetFollowResponse, 0, len(follows))
	for _, follow := range follows {
		followResponses = append(followResponses, &GetFollowResponse{
			UserID:   follow.GetPreferredUserID(),
			Username: follow.GetPreferredUserName(),
			Summary:  "", // TODO: Return actual value https://github.com/stlatica/stlatica/issues/524
			Icon:     "", // TODO: Return actual value https://github.com/stlatica/stlatica/issues/524
			IsPublic: follow.GetIsPublic(),
		})
	}
	return followResponses, nil
}

func (c *userController) PostFollow(ectx echo.Context, userIDStr string, followUserIDStr string) (error) {
	user, err := c.userUseCase.GetUserByPreferredUserID(ectx.Request().Context(), userIDStr)
	if err != nil {
		return err
	}
	followUser, err := c.userUseCase.GetUserByPreferredUserID(ectx.Request().Context(), followUserIDStr)
	if err != nil {
		return err
	}
	return c.userUseCase.FollowUser(user, followUser)
}


func (c *userController) GetFollowers(ectx echo.Context,
	userIDStr string, userPaginationID *string, limit *int) ([]*GetFollowerResponse, error) {
	var limitValue int
	var userPaginationIDStr string
	if userPaginationID == nil {
		userPaginationIDStr = ""
	} else {
		userPaginationIDStr = *userPaginationID
	}
	if limit != nil {
		limitValue = *limit
	}

	getParams := ports.FollowersGetParams{
		PreferredUserID:           userIDStr,
		PreferredUserPaginationID: userPaginationIDStr,
		Limit:                     limitValue,
	}
	followers, err := c.userUseCase.GetFollowers(ectx.Request().Context(), getParams)
	if err != nil {
		return nil, err
	}

	followerResponse := make([]*GetFollowerResponse, 0, len(followers))
	for _, follower := range followers {
		followerResponse = append(followerResponse, &GetFollowerResponse{
			UserID:   follower.GetPreferredUserID(),
			Username: follower.GetPreferredUserName(),
			Summary:  "", // TODO: Return actual value https://github.com/stlatica/stlatica/issues/526
			Icon:     "", // TODO: Return actual value https://github.com/stlatica/stlatica/issues/526
			IsPublic: follower.GetIsPublic(),
		})
	}
	return followerResponse, nil
}

func (c *userController) GetUserIcon(ectx echo.Context, preferredUserIDStr string) (io.ReadCloser, error) {
	return c.userUseCase.GetUserIcon(ectx.Request().Context(), preferredUserIDStr)
}
