package interfaces

import "github.com/athunlal/bookNowTrain-svc/pkg/domain"

type TrainUseCase interface {
	AddTrain(train domain.Train) error
}
