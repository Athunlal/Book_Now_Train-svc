package usecase

import (
	"context"
	"errors"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/repository/interface"
	usecase "github.com/athunlal/bookNowTrain-svc/pkg/usecase/interface"
)

type TrainUseCase struct {
	Repo interfaces.TrainRepo
}

// UpdateTrainRoute implements interfaces.TrainUseCase.
func (use *TrainUseCase) UpdateTrainRoute(ctx context.Context, trainData domain.Train) error {
	err := use.Repo.UpdateTrainRoute(ctx, trainData)
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
