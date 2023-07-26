package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Train struct {
	TrainId     uint   `json:"trainid" bson:"trainid,omitempty"`
	TrainNumber uint   `json:"trainumber" bson:"trainumber,omitempty" validate:"required,min=4,max=50"`
	TrainName   string `json:"trainname" bson:"trainname,omitempty" validate:"required,min=2,max=50"`
	Route       []uint `json:"route,omitempty" bson:"route,omitempty"`
}

type Station struct {
	StationId   uint   `json:"stationid" bson:"_id,omitempty"`
	StationName string `json:"stationname" bson:"stationname,omitempty"`
	City        string `json:"city" bson:"city,omitempty"`
}
type Route struct {
	RouteId  primitive.ObjectID `json:"routeid" bson:"_id,omitempty"`
	RouteMap []RouteStation     `json:"routemap,omitempty" bson:"routemap,omitempty"`
}

type RouteStation struct {
	StationId primitive.ObjectID `json:"stationid" bson:"stationid,omitempty"`
	Distance  float64            `json:"distance,omitempty" bson:"distance,omitempty"`
	Time      string             `json:"time,omitempty" bson:"time,omitempty"`
}
