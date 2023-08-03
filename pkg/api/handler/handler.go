package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	"github.com/athunlal/bookNowTrain-svc/pkg/pb"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/usecase/interface"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TrainHandler struct {
	useCase interfaces.TrainUseCase
	pb.TrainManagementServer
}

func NewTrainHandler(usecase interfaces.TrainUseCase) *TrainHandler {
	return &TrainHandler{
		useCase: usecase,
	}
}

func (h *TrainHandler) AddRoute(ctx context.Context, req *pb.AddRouteRequest) (*pb.AddRouteResponse, error) {

	routeid, err := primitive.ObjectIDFromHex(req.Route.Routeid)
	if err != nil {
		log.Fatal("Converting the string to primitive.ObjectId err", err)
	}

	routeData := domain.Route{
		RouteName: req.Route.Routename,
		RouteId:   routeid,
		RouteMap:  make([]domain.RouteStation, len(req.Route.Routemap)),
	}

	for i, rs := range req.Route.Routemap {
		stationID, err := primitive.ObjectIDFromHex(rs.Stationid)
		if err != nil {
			log.Fatal("Converting the string to primitive.ObjectID err", err)
		}
		routeData.RouteMap[i] = domain.RouteStation{
			Distance:  rs.Distance,
			Time:      rs.Time,
			StationId: stationID,
		}
	}

	err = h.useCase.AddRoute(ctx, routeData)

	response := &pb.AddRouteResponse{
		// Set the appropriate fields based on the processing result.
		// For example, you can set status and error fields if needed.
	}

	return response, nil

}

func (h *TrainHandler) AddStation(ctx context.Context, req *pb.AddStationRequest) (*pb.AddStationResponse, error) {
	station := domain.Station{
		StationName: req.Stationname,
		City:        req.City,
	}

	_, err := h.useCase.AddStation(ctx, station)
	if err != nil {
		return &pb.AddStationResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}

	return &pb.AddStationResponse{
		Status: http.StatusOK,
	}, nil
}

func (h *TrainHandler) AddTrain(ctx context.Context, req *pb.AddTrainRequest) (*pb.AddTrainResponse, error) {
	train := domain.Train{
		TrainName:   req.Trainname,
		TrainNumber: uint(req.Trainnumber),
	}

	err := h.useCase.AddTrain(ctx, train)
	if err != nil {
		return &pb.AddTrainResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}

	return &pb.AddTrainResponse{
		Status: http.StatusOK,
	}, nil
}
