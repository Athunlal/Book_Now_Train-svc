package interfaces

import (
	"context"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
)

type SationUseCase interface {
	AddStation(ctx context.Context, station domain.Station) error
}
