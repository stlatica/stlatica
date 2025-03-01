package v1

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/stlatica/stlatica/packages/backend/app/controllers/internalapi/v1/openapi"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

const maxPlatNum = 10

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(server *echo.Echo) {
	openapi.RegisterHandlers(server, &handler{})
}

type handler struct{}

// DeleteImage implements openapi.ServerInterface.
func (h *handler) DeleteImage(_ echo.Context, _ string) error {
	panic("unimplemented")
}

func (h *handler) Login(_ echo.Context) error {
	panic("implement me")
}

func (h *handler) DeleteFavorite(_ echo.Context, _ openapi.PlatId) error {
	panic("implement me")
}

func (h *handler) PostFavorite(_ echo.Context, _ openapi.PlatId) error {
	panic("implement me")
}

func (h *handler) GetUsers(_ echo.Context, _ openapi.GetUsersParams) error {
	panic("implement me")
}

func (h *handler) CreateUser(_ echo.Context) error {
	panic("implement me")
}

func (h *handler) UpdateUser(_ echo.Context, _ openapi.UserId) error {
	panic("implement me")
}

func (h *handler) DeleteUser(_ echo.Context, _ openapi.UserId) error {
	panic("implement me")
}

func (h *handler) PostPlat(_ echo.Context) error {
	panic("implement me")
}

func (h *handler) DeletePlat(_ echo.Context, _ string) error {
	panic("implement me")
}

func (h *handler) GetPlat(_ echo.Context, _ string) error {
	panic("implement me")
}

func (h *handler) GetTimeline(ectx echo.Context, _ string) error {
	var plats []openapi.Plat
	now := time.Now()
	platNum, err := rand.Int(rand.Reader, big.NewInt(maxPlatNum))
	if err != nil {
		panic("failed to create num")
	}
	if platNum.Int64() == 0 {
		return ectx.JSON(http.StatusOK, []openapi.Plat{})
	}
	for i := int(platNum.Int64()); i >= 0; i-- {
		plats = append(plats, openapi.Plat{
			PlatId: ulid.MustNew(ulid.Timestamp(
				types.NewUnixTimeFromTime(now.Add(time.Duration(i)*time.Minute)).ConvertToTime()),
				rand.Reader,
			).String(),
			Content: fmt.Sprintf("plat %d", i),
			ImageUrls: &[]string{
				"localhost:8080/internal/v1/images/1",
				"localhost:8080/internal/v1/images/2",
				"localhost:8080/internal/v1/images/3",
			},
			CreatedAt: now,
		})
	}
	return ectx.JSON(http.StatusOK, plats)
}

func (h *handler) GetTimelineByQuery(_ echo.Context, _ openapi.GetTimelineByQueryParams) error {
	panic("implement me")
}

// GetImage is the handler for GET /images/{image_id}, ServerInterface implementation.
// if image_id is "large", return large image, otherwise return small image.
func (h *handler) GetImage(ectx echo.Context, imageID string) error {
	if imageID == "large" {
		return ectx.String(http.StatusOK, mockLargeImage)
	}
	return ectx.String(http.StatusOK, mockSmallImage)
}

func (h *handler) UploadImage(_ echo.Context) error {
	panic("implement me")
}

func (h *handler) GetUser(_ echo.Context, _ string) error {
	panic("implement me")
}

func (h *handler) GetFollows(_ echo.Context, _ openapi.UserId, _ openapi.GetFollowsParams) error {
	panic("implement me")
}

func (h *handler) DeleteFollow(_ echo.Context, _ openapi.UserId) error {
	panic("implement me")
}

func (h *handler) PostFollow(_ echo.Context, _ openapi.UserId) error {
	panic("implement me")
}

func (h *handler) GetFollowers(_ echo.Context, _ openapi.UserId, _ openapi.GetFollowersParams) error {
	panic("implement me")
}

func (h *handler) GetUserIcon(_ echo.Context, _ openapi.UserId) error {
	panic("implement me")
}
