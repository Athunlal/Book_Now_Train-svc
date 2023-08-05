package interfaces

import (
	"context"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
)

type TrainUseCase interface {
	AddTrain(ctx context.Context, train domain.Train) error
	AddStation(ctx context.Context, station domain.Station) (domain.Station, error)
	AddRoute(ctx context.Context, route domain.Route) error
	UpdateTrainRoute(ctx context.Context, trainData domain.Train) error
	SearchTrain(ctx context.Context, searcheData domain.SearchingTrainRequstedData) (domain.SearchingTrainResponseData, error)

	ViewTrain(ctx context.Context) (*domain.SearchingTrainResponseData, error)
}
