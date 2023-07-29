package interfaces

import (
	"context"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
)

type TrainUseCase interface {
	AddTrain(ctx context.Context, train domain.Train) error
}
