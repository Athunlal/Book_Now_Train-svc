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
