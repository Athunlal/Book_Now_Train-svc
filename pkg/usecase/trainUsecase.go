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

// AddRoute implements interfaces.TrainUseCase.
func (use *TrainUseCase) AddRoute(ctx context.Context, route domain.Route) error {
	err := use.Repo.AddRoute(ctx, route)
	if err != nil {
		return errors.New(" name is already exist")
	}
	return err
}

// AddStation implements interfaces.TrainUseCase.
func (use *TrainUseCase) AddStation(ctx context.Context, station domain.Station) error {
	_, err := use.Repo.FindByStationName(ctx, station)
	if err == nil {
		return errors.New("Station name is already exist")
	}
	_, err = use.Repo.FindByStationid(ctx, station)
	if err == nil {
		return errors.New("Station id  is already exist")
	}
	err = use.Repo.AddStation(ctx, station)

	return err
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
