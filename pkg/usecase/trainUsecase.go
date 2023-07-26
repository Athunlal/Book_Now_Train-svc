package usecase

import (
	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/repository/interface"
	usecase "github.com/athunlal/bookNowTrain-svc/pkg/usecase/interface"
)

type TrainUseCase struct {
	Repo interfaces.TrainRepo
}

func (use *TrainUseCase) AddTrain(trainData domain.Train) error {
	err := use.Repo.AddTrain(trainData)
	return err
}

func NewTrainUseCase(repo interfaces.TrainRepo) usecase.TrainUseCase {
	return &TrainUseCase{
		Repo: repo,
	}
}
