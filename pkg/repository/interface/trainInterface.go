package interfaces

import (
	"context"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
)

type TrainRepo interface {
	FindbyTrainName(ctx context.Context, train domain.Train) (domain.Train, error)
	FindByTrainNumber(tx context.Context, train domain.Train) (domain.Train, error)
	AddTrain(tx context.Context, train domain.Train) error

	FindByStationName(ctx context.Context, station domain.Station) (domain.Station, error)
	AddStation(ctx context.Context, station domain.Station) error
	FindroutebyName(ctx context.Context, route domain.Route) (domain.Route, error)
	AddRoute(ctx context.Context, route domain.Route) error

	UpdateTrainRoute(ctx context.Context, trainData domain.Train) error
}
