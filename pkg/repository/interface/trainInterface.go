package interfaces

import "github.com/athunlal/bookNowTrain-svc/pkg/domain"

type TrainRepo interface {
	AddTrain(train domain.Train) error
}
