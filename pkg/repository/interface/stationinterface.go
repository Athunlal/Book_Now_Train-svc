package interfaces

import (
	"context"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
)

type SationRepo interface {
	FindByStationName(ctx context.Context, station domain.Station) (domain.Station, error)
	FindByStationid(ctx context.Context, station domain.Station) (domain.Station, error)
	AddStation(ctx context.Context, station domain.Station) error
}
