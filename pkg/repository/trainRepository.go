package repository

import (
	"context"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/mongo"
)

type TrainDataBase struct {
	DB *mongo.Database
}

func (db *TrainDataBase) AddTrain(train domain.Train) error {
	collection := db.DB.Collection("station")
	_, err := collection.InsertOne(context.Background(), train)
	return err
}

// AddTrain implements interfaces.TrainRepo.

func NewTrainRepo(db *mongo.Database) interfaces.TrainRepo {
	return &TrainDataBase{
		DB: db,
	}
}
