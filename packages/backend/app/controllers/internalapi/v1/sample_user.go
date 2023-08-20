package v1

type sampleUserController struct{}

type GetSampleUserResponse struct {
	UserID   string
	Username string
}

// GetSampleUser converts request data and calls usecase to get sample user.
func (c *sampleUserController) GetSampleUser(sampleUserID string) (GetSampleUserResponse, error) {
	// TODO: add usecase call https://github.com/stlatica/stlatica/issues/51
	return GetSampleUserResponse{
		UserID:   sampleUserID,
		Username: "sample_user",
	}, nil
}
