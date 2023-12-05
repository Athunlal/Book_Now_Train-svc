package interfaces

import (
	"context"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	SearchTrain(ctx context.Context, searcheData domain.SearchingTrainRequstedData) (domain.SearchingTrainResponseData, error)
	SearchTrainbyName(ctx context.Context, tran_name string) (domain.Train, error)

	ViewTrain(ctx context.Context) (*domain.SearchingTrainResponseData, error)
	ViewStation(ctx context.Context) (*domain.SearchStationRes, error)

	AddSeat(ctc context.Context, seat domain.Seats) (error, *mongo.InsertOneResult)
	FindSeatbyCompartment(ctx context.Context, seat domain.Seats) (error, domain.Seats)
	UpdateSeatIntoTrain(ctx context.Context, updateData domain.Train) error

	FindCompartmentByid(ctx context.Context, compartmentId primitive.ObjectID) error
	ViewRoute(ctx context.Context) (*[]domain.Route, error)
}
