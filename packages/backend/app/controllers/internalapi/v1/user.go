package v1

type userController struct{}

type GetUserResponse struct{}

// GetUser converts request data and calls usecase to get user.
func (c *userController) GetUser(_ string) (GetUserResponse, error) {
	// TODO: add usecase call https://github.com/stlatica/stlatica/issues/167
	return GetUserResponse{}, nil
}
