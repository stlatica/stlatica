package plats

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/plats"
	domainports "github.com/stlatica/stlatica/packages/backend/app/domains/plats/ports"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/plats/ports"
)

// PlatUseCase is the interface for getting plat.
type PlatUseCase interface {
	// GetPlat returns plat.
	GetPlat(ctx context.Context, platID types.PlatID) (*entities.Plat, error)
	// GetPlats returns plats.
	GetPlats(ctx context.Context, userID types.UserID, getParams ports.PlatsGetParams) ([]*entities.Plat, error)
	// CreatePlat creates new plat.
	CreatePlat(ctx context.Context, userID types.UserID, content string) (*entities.Plat, error)
}

// NewPlatUseCase returns PlatUseCase.
func NewPlatUseCase(appLogger *logger.AppLogger, domainFactory plats.Factory, platDAO dao.PlatDAO) PlatUseCase {
	return &platUseCase{
		appLogger:     appLogger,
		platDAO:       platDAO,
		domainFactory: domainFactory,
	}
}

type platUseCase struct {
	appLogger     *logger.AppLogger
	platDAO       dao.PlatDAO
	domainFactory plats.Factory
}

func (u *platUseCase) GetPlat(ctx context.Context, platID types.PlatID) (*entities.Plat, error) {
	getter := u.domainFactory.NewPlatGetter()
	portImpl := &platPortImpl{
		platDAO: u.platDAO,
	}
	return getter.GetPlat(ctx, platID, portImpl)
}

func (u *platUseCase) GetPlats(ctx context.Context,
	userID types.UserID, getParams ports.PlatsGetParams) ([]*entities.Plat, error) {
	getter := u.domainFactory.NewPlatGetter()
	portImpl := &platPortImpl{
		platDAO: u.platDAO,
	}

	var userIDs []*types.UserID
	switch getParams.TimelineType {
	case "home":
		userIDs = []*types.UserID{
			&userID,
		}
	case "following":
		// TODO: timeline_type == following のケース実装 https://github.com/stlatica/stlatica/issues/440
		panic("implement me")
	case "local":
		// TODO: timeline_type == local のケース実装 https://github.com/stlatica/stlatica/issues/440
		panic("implement me")
	default:
		return nil, errors.New("invalid timeline_type")
	}

	domainParams := domainports.PlatsGetParams{
		UserIDs:  userIDs,
		ToDate:   getParams.ToDate,
		FromDate: getParams.FromDate,
		Limit:    getParams.Limit,
	}
	return getter.GetPlats(ctx, domainParams, portImpl)
}

func (u *platUseCase) CreatePlat(ctx context.Context, userID types.UserID, content string) (*entities.Plat, error) {
	creator := u.domainFactory.NewPlatCreator()
	portImpl := &platPortImpl{
		platDAO: u.platDAO,
	}
	return creator.CreatePlat(ctx, userID, content, portImpl)
}

type platPortImpl struct {
	platDAO dao.PlatDAO
}

func (p *platPortImpl) GetPlat(ctx context.Context, platID types.PlatID) (*entities.Plat, error) {
	return p.platDAO.GetPlat(ctx, platID)
}

func (p *platPortImpl) GetPlats(ctx context.Context,
	getParams domainports.PlatsGetParams) ([]*entities.Plat, error) {
	return p.platDAO.GetPlats(ctx, getParams)
}

func (p *platPortImpl) CreatePlat(ctx context.Context, userID types.UserID, content string) (*entities.Plat, error) {
	return p.platDAO.CreatePlat(ctx, userID, content)
}
