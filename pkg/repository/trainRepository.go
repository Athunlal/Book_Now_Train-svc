package repository

import (
	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/mongo"
)

type TrainDataBase struct {
	DB *mongo.Database
}

// AddTrain implements interfaces.TrainRepo.
func (*TrainDataBase) AddTrain(train domain.Train) error {
	panic("unimplemented")
}

// AddTrain implements interfaces.TrainRepo.

func NewTrainRepo(db *mongo.Database) interfaces.TrainRepo {
	return &TrainDataBase{
		DB: db,
	}
}
