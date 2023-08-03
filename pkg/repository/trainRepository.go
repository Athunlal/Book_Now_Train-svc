package repository

import (
	"context"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TrainDataBase struct {
	DB *mongo.Database
}

// FindroutebyName implements interfaces.TrainRepo.
func (db *TrainDataBase) FindroutebyName(ctx context.Context, route domain.Route) (domain.Route, error) {
	filter := bson.M{"": route.RouteName}
	var result domain.Route
	err := db.DB.Collection("routename").FindOne(ctx, filter).Decode(&result)
	return result, err
}

// AddRoute implements interfaces.TrainRepo.
func (db *TrainDataBase) AddRoute(ctx context.Context, route domain.Route) error {
	collection := db.DB.Collection("route")
	_, err := collection.InsertOne(context.Background(), route)
	return err
}

// AddRoute implements interfaces.TrainRepo.

// AddStation implements interfaces.TrainRepo.
func (db *TrainDataBase) AddStation(ctx context.Context, station domain.Station) error {
	collection := db.DB.Collection("station")
	_, err := collection.InsertOne(ctx, station)
	return err
}

// FindByStationName implements interfaces.TrainRepo.
func (db *TrainDataBase) FindByStationName(ctx context.Context, station domain.Station) (domain.Station, error) {
	filter := bson.M{"stationname": station.StationName}
	var result domain.Station
	err := db.DB.Collection("station").FindOne(ctx, filter).Decode(&result)
	return result, err
}

// FindbyTrainName implements interfaces.TrainRepo.
func (db *TrainDataBase) FindbyTrainName(ctx context.Context, train domain.Train) (domain.Train, error) {
	filter := bson.M{"trainname": train.TrainName}
	var result domain.Train
	err := db.DB.Collection("train").FindOne(ctx, filter).Decode(&result)

	return result, err
}

// AddTrain implements interfaces.TrainRepo.
func (db *TrainDataBase) AddTrain(tx context.Context, train domain.Train) error {
	collection := db.DB.Collection("train")
	_, err := collection.InsertOne(tx, train)
	return err

}

// FindByTrainNumber implements interfaces.TrainRepo.
func (db *TrainDataBase) FindByTrainNumber(tx context.Context, train domain.Train) (domain.Train, error) {
	filter := bson.M{"trainumber": train.TrainNumber}
	var result domain.Train
	err := db.DB.Collection("train").FindOne(tx, filter).Decode(&result)

	return result, err

}

// AddTrain implements interfaces.TrainRepo.
func NewTrainRepo(db *mongo.Database) interfaces.TrainRepo {
	return &TrainDataBase{
		DB: db,
	}
}
