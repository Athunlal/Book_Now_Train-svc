package interfaces

import (
	"context"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type TrainUseCase interface {
	AddTrain(ctx context.Context, train domain.Train) error
	AddStation(ctx context.Context, station domain.Station) (domain.Station, error)
	AddRoute(ctx context.Context, route domain.Route) error
	UpdateTrainRoute(ctx context.Context, trainData domain.Train) error
	SearchTrain(ctx context.Context, searcheData domain.SearchingTrainRequstedData) (domain.SearchingTrainResponseData, error)
	AddSeat(ctx context.Context, seat domain.SeatData) (error, *mongo.InsertOneResult)
	ViewTrain(ctx context.Context) (*domain.SearchingTrainResponseData, error)
	UpadateSeatInotTrain(ctx context.Context, updateData domain.Train) error
}
