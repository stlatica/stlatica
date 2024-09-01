package v1

import (
	"github.com/labstack/echo/v4"
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

// GetFollowResponse is the response of GetFollows.
type GetFollowResponse struct {
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
func (c *userController) CreateUser(ectx echo.Context, userName string, mailAddress string) (
	*GetUserResponse, error) {
	user, err := c.userUseCase.CreateUser(ectx.Request().Context(), userName, mailAddress)
	if err != nil {
		return nil, err
	}
	return &GetUserResponse{
		UserID: user.GetPreferredUserID(),
	}, nil
}

func (c *userController) GetFollows(ectx echo.Context,
	userIDStr string, userPaginationID *string, limit *int) ([]*GetFollowResponse, error) {
	var limitValue uint64
	var userPaginationIDStr string
	if limit != nil {
		limitValue = uint64(*limit)
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
