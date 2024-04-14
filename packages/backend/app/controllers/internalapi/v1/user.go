package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/users"
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
