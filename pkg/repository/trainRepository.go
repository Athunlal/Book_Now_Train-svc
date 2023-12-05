package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TrainDataBase struct {
	DB *mongo.Database
}

// ViewRoute implements interfaces.TrainRepo.
func (db *TrainDataBase) ViewRoute(ctx context.Context) (*[]domain.Route, error) {
	var routes []domain.Route
	cursor, err := db.DB.Collection("route").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var route domain.Route
		if err := cursor.Decode(&route); err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &routes, nil
}

// FindCompartmentByid implements interfaces.TrainRepo.
func (db *TrainDataBase) FindCompartmentByid(ctx context.Context, compartmentId primitive.ObjectID) error {
	collectionSeat := db.DB.Collection("seat")
	filter := bson.M{"_id": compartmentId}
	result := collectionSeat.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

// SearchTrainby_name implements interfaces.TrainRepo.
func (db *TrainDataBase) SearchTrainbyName(ctx context.Context, trainName string) (domain.Train, error) {
	collectionTrain := db.DB.Collection("train")

	filter := bson.M{"trainName": trainName}

	result := collectionTrain.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		return domain.Train{}, err
	}

	var train domain.Train
	if err := result.Decode(&train); err != nil {
		log.Printf("Error decoding document: %v\n", err)
		return domain.Train{}, err
	}

	return train, nil
}

// ViewStation implements interfaces.TrainRepo.
func (db *TrainDataBase) ViewStation(ctx context.Context) (*domain.SearchStationRes, error) {
	var Station domain.SearchStationRes
	cursor, err := db.DB.Collection("station").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var station domain.Station
		if err := cursor.Decode(&station); err != nil {
			return nil, err
		}
		Station.Station = append(Station.Station, station)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &Station, nil
}

// UpdateSeatIntoTrain implements interfaces.TrainRepo.
func (db *TrainDataBase) UpdateSeatIntoTrain(ctx context.Context, updateData domain.Train) error {
	collection := db.DB.Collection("train")
	filter := bson.M{"trainNumber": updateData.TrainNumber}
	update := bson.M{"$set": bson.M{"compartment": updateData.Compartment}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return err
}

// FindSeatbyCompartment implements interfaces.TrainRepo.
func (db *TrainDataBase) FindSeatbyCompartment(ctx context.Context, seat domain.Seats) (error, domain.Seats) {
	filter := bson.M{"compartment": seat.Compartment}
	var result domain.Seats
	err := db.DB.Collection("seat").FindOne(ctx, filter).Decode(&result)
	return err, result
}

// AddSeat implements interfaces.TrainRepo.
func (db *TrainDataBase) AddSeat(ctx context.Context, seat domain.Seats) (error, *mongo.InsertOneResult) {
	collection := db.DB.Collection("seat")

	res, err := collection.InsertOne(ctx, seat)
	if err != nil {
		return err, nil
	}

	return nil, res
}

// ViewTrain implements interfaces.TrainRepo.
func (db *TrainDataBase) ViewTrain(ctx context.Context) (*domain.SearchingTrainResponseData, error) {
	var Train []domain.Train
	cursor, err := db.DB.Collection("train").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var train domain.Train
		if err := cursor.Decode(&train); err != nil {
			return nil, err
		}
		Train = append(Train, train)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &domain.SearchingTrainResponseData{
		SearcheResponse: Train,
	}, nil
}

func (db *TrainDataBase) SearchTrain(ctx context.Context, searchData domain.SearchingTrainRequstedData) (domain.SearchingTrainResponseData, error) {
	collectionTrain := db.DB.Collection("train")

	sourceStationID := searchData.SourceStationid
	destinationStationID := searchData.DestinationStationid

	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.M{
			"from":         "route",
			"localField":   "route",
			"foreignField": "_id",
			"as":           "train_route",
		}}},
		{{Key: "$unwind", Value: "$train_route"}},
		{{Key: "$match", Value: bson.M{
			"train_route.routemap.stationid": bson.M{"$all": []primitive.ObjectID{sourceStationID, destinationStationID}},
		}}},
	}

	cur, err := collectionTrain.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Printf("Error executing aggregation pipeline: %v\n", err)
		return domain.SearchingTrainResponseData{}, err
	}
	defer cur.Close(context.Background())

	var trains []domain.Train
	for cur.Next(context.Background()) {
		var train domain.Train
		if err := cur.Decode(&train); err != nil {
			log.Printf("Error decoding document: %v\n", err)
			return domain.SearchingTrainResponseData{}, err
		}
		trains = append(trains, train)
	}

	if err := cur.Err(); err != nil {
		log.Printf("Error reading cursor: %v\n", err)
		return domain.SearchingTrainResponseData{}, err
	}

	if len(trains) == 0 {
		return domain.SearchingTrainResponseData{
			SearcheResponse: nil,
		}, errors.New("Train not found")
	}

	return domain.SearchingTrainResponseData{
		SearcheResponse: trains,
	}, nil
}

// UpdateTrainRoute implements interfaces.TrainRepo.
func (db *TrainDataBase) UpdateTrainRoute(ctx context.Context, trainData domain.Train) error {
	collection := db.DB.Collection("train")
	filter := bson.M{"trainNumber": trainData.TrainNumber}
	update := bson.M{"$set": bson.M{"route": trainData.Route}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("Error updating train: %v\n", err)
	}
	return err
}

// FindroutebyName implements interfaces.TrainRepo.
func (db *TrainDataBase) FindroutebyName(ctx context.Context, route domain.Route) (domain.Route, error) {
	filter := bson.M{"routename": route.RouteName}
	var result domain.Route
	err := db.DB.Collection("route").FindOne(ctx, filter).Decode(&result)
	return result, err
}

// AddRoute implements interfaces.TrainRepo.
func (db *TrainDataBase) AddRoute(ctx context.Context, route domain.Route) error {
	collection := db.DB.Collection("route")
	_, err := collection.InsertOne(context.Background(), route)
	return err
}

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
	filter := bson.M{"trainName": train.TrainName}
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
	filter := bson.M{"trainNumber": train.TrainNumber}
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
