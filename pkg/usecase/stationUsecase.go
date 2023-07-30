package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/repository/interface"
	usecase "github.com/athunlal/bookNowTrain-svc/pkg/usecase/interface"
)

type SationUseCase struct {
	Repo interfaces.SationRepo
}

// AddStation implements interfaces.SationUseCase.
func (use *SationUseCase) AddStation(ctx context.Context, station domain.Station) error {
	fmt.Println("this sis stattion function ----------->>>>>>>>>>")
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

func NewSationUseCase(repo interfaces.SationRepo) usecase.SationUseCase {
	return &SationUseCase{
		Repo: repo,
	}
}
