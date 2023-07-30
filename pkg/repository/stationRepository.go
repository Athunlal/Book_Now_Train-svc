package repository

import (
	"context"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SationDataBase struct {
	DB *mongo.Database
}

// AddStation implements interfaces.SationRepo.
func (db *SationDataBase) AddStation(ctx context.Context, station domain.Station) error {
	collection := db.DB.Collection("station")
	_, err := collection.InsertOne(ctx, station)
	return err

}

// FindByStationName implements interfaces.SationRepo.
func (db *SationDataBase) FindByStationName(ctx context.Context, station domain.Station) (domain.Station, error) {
	filter := bson.M{"trainname": station.StationName}
	var result domain.Station
	err := db.DB.Collection("station").FindOne(ctx, filter).Decode(&result)
	return result, err
}

// FindByStationid implements interfaces.SationRepo.
func (db *SationDataBase) FindByStationid(ctx context.Context, station domain.Station) (domain.Station, error) {
	filter := bson.M{"trainname": station.StationId}
	var result domain.Station
	err := db.DB.Collection("train").FindOne(ctx, filter).Decode(&result)
	return result, err
}

// AddTrain implements interfaces.TrainRepo.
func NewSationRepo(db *mongo.Database) interfaces.SationRepo {
	return &SationDataBase{
		DB: db,
	}
}
