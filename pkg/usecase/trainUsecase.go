package usecase

import (
	"context"
	"errors"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/repository/interface"
	usecase "github.com/athunlal/bookNowTrain-svc/pkg/usecase/interface"
	"github.com/athunlal/bookNowTrain-svc/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type TrainUseCase struct {
	Repo interfaces.TrainRepo
}

// UpadateSeatInotTrain implements interfaces.TrainUseCase.
func (use *TrainUseCase) UpadateSeatInotTrain(ctx context.Context, updateData domain.Train) error {
	_, err := use.Repo.FindByTrainNumber(ctx, updateData)
	if err != nil {
		return errors.New("Train number not exist")
	}
	err = use.Repo.UpdateSeatIntoTrain(ctx, updateData)
	return err
}

// AddSeat implements interfaces.TrainUseCase.
func (use *TrainUseCase) AddSeat(ctx context.Context, seat domain.SeatData) (error, *mongo.InsertOneResult) {
	allocatedSeate := utils.SeateAllocation(seat)
	err, response := use.Repo.FindSeatbyCompartment(ctx, allocatedSeate)
	var res *mongo.InsertOneResult
	if err != nil {
		err, res = use.Repo.AddSeat(ctx, allocatedSeate)
		return err, res
	}
	if response.Compartment == seat.Compartment {
		return errors.New("Compartment name is already exist"), nil
	}
	return nil, res
}

// ViewTrain implements interfaces.TrainUseCase.
func (use *TrainUseCase) ViewTrain(ctx context.Context) (*domain.SearchingTrainResponseData, error) {
	res, err := use.Repo.ViewTrain(ctx)
	return res, err
}

// SearchTrain implements interfaces.TrainUseCase.
func (use *TrainUseCase) SearchTrain(ctx context.Context, searcheData domain.SearchingTrainRequstedData) (domain.SearchingTrainResponseData, error) {
	res, err := use.Repo.SearchTrain(ctx, searcheData)
	return res, err
}

// UpdateTrainRoute implements interfaces.TrainUseCase.
func (use *TrainUseCase) UpdateTrainRoute(ctx context.Context, trainData domain.Train) error {
	_, err := use.Repo.FindByTrainNumber(ctx, trainData)
	if err == nil {
		err = use.Repo.UpdateTrainRoute(ctx, trainData)
	} else {
		return errors.New("Invalid train number")
	}
	return err
}

// AddRoute implements interfaces.TrainUseCase.
func (use *TrainUseCase) AddRoute(ctx context.Context, route domain.Route) error {
	result, err := use.Repo.FindroutebyName(ctx, route)
	if result.RouteName == "" {
		err = use.Repo.AddRoute(ctx, route)
	} else {
		return errors.New("Route name  is already exist")
	}
	return err
}

// AddStation implements interfaces.TrainUseCase.
func (use *TrainUseCase) AddStation(ctx context.Context, station domain.Station) (domain.Station, error) {
	result, err := use.Repo.FindByStationName(ctx, station)
	if result.StationName == "" {
		err = use.Repo.AddStation(ctx, station)
	} else {
		return result, errors.New("station name  is already exist")
	}

	return result, err
}

// AddTrain implements interfaces.TrainUseCase.
func (use *TrainUseCase) AddTrain(ctx context.Context, train domain.Train) error {
	_, err := use.Repo.FindByTrainNumber(ctx, train)
	if err == nil {
		return errors.New("Train number is already exist")
	}
	_, err = use.Repo.FindbyTrainName(ctx, train)
	if err == nil {
		return errors.New("Train name is already exist")
	}
	err = use.Repo.AddTrain(ctx, train)

	return err
}

func NewTrainUseCase(repo interfaces.TrainRepo) usecase.TrainUseCase {
	return &TrainUseCase{
		Repo: repo,
	}
}
