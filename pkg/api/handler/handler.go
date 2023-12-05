package handler

import (
	"context"
	"fmt"
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

func (h *TrainHandler) ViewStation(ctx context.Context, req *pb.ViewRequest) (*pb.ViewStationResponse, error) {
	res, err := h.useCase.ViewStation(ctx)

	if err != nil {
		return &pb.ViewStationResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}

	// Create a slice to store the converted Station data
	stations := make([]*pb.Station, len(res.Station))

	// Convert domain.Station instances to pb.Station instances
	for i, station := range res.Station {
		pbStation := &pb.Station{
			Stationid:   station.StationId.Hex(),
			StationName: station.StationName,
			City:        station.City,
		}
		stations[i] = pbStation
	}

	// Create the response with the converted station data
	stationResponse := &pb.ViewStationResponse{
		Status:   http.StatusOK,
		Stations: stations,
	}

	return stationResponse, nil
}

func (h *TrainHandler) UpdateSeatIntoTrain(ctx context.Context, req *pb.UpdateSeatIntoTrainRequest) (*pb.UpdateSeatIntoTrainResponse, error) {
	fmt.Println("This is the trian number : ", req.Trainnumber)

	updateData := domain.Train{
		TrainNumber: uint(req.Trainnumber),
		Compartment: make([]domain.Compartment, len(req.Compartments)),
	}

	for i, rs := range req.Compartments {
		seatid, _ := primitive.ObjectIDFromHex(rs.Seatid)
		updateData.Compartment[i] = domain.Compartment{
			Seatid: seatid,
		}
	}

	err := h.useCase.UpadateSeatInotTrain(ctx, updateData)
	if err != nil {
		return &pb.UpdateSeatIntoTrainResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}
	return &pb.UpdateSeatIntoTrainResponse{
		Status: http.StatusOK,
	}, nil
}

func (h *TrainHandler) AddSeat(ctx context.Context, req *pb.AddSeatRequest) (*pb.AddSeatResponse, error) {
	seatdata := domain.SeatData{
		Price:         float32(req.Price),
		NumbserOfSeat: int(req.Numberofseat),
		TypeOfSeat:    req.Typeofseat,
		Compartment:   req.Compartment,
	}
	err, _ := h.useCase.AddSeat(ctx, seatdata)

	if err != nil {
		return &pb.AddSeatResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}
	return &pb.AddSeatResponse{
		Status: http.StatusOK,
	}, nil
}

func (h *TrainHandler) ViewTrain(ctx context.Context, req *pb.ViewTrainRequest) (*pb.ViewTrainResponse, error) {
	res, err := h.useCase.ViewTrain(ctx)
	if err != nil {
		return &pb.ViewTrainResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}

	response := &pb.ViewTrainResponse{
		Traindata: make([]*pb.Train, len(res.SearcheResponse)),
		Status:    http.StatusOK,
	}
	for i, rs := range res.SearcheResponse {
		response.Traindata[i] = &pb.Train{
			Trainname:   rs.TrainName,
			Trainnumber: int64(rs.TrainNumber),
		}
	}
	return response, nil
}

func (h *TrainHandler) UpdateTrainRoute(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	routeid, err := primitive.ObjectIDFromHex(req.Route)
	if err != nil {
		log.Fatal("Converting the string to primitive.ObjectId err", err)
	}

	trainData := domain.Train{
		Route:       routeid,
		TrainNumber: uint(req.Trainnumber),
	}

	err = h.useCase.UpdateTrainRoute(ctx, trainData)
	if err != nil {
		return &pb.UpdateResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}

	return &pb.UpdateResponse{
		Status: http.StatusOK,
	}, nil
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
	if err != nil {
		return &pb.AddRouteResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error Found in usecase",
		}, err
	}

	return &pb.AddRouteResponse{
		Status: http.StatusOK,
	}, nil

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
		TrainName:      req.Trainname,
		TrainNumber:    uint(req.Trainnumber),
		TrainType:      req.Traintype,
		StartingTime:   req.Startingtime,
		EndingtingTime: req.Endingtime,
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

func (h *TrainHandler) SearchTrain(ctx context.Context, req *pb.SearchTrainRequest) (*pb.SearchTrainResponse, error) {

	sourceid, err := primitive.ObjectIDFromHex(req.Sourcestationid)
	if err != nil {
		log.Fatal("Converting the string to primitive.ObjectId err", err)
	}

	destinationid, err := primitive.ObjectIDFromHex(req.Destinationstationid)
	if err != nil {
		log.Fatal("Converting the string to primitive.ObjectId err", err)
	}

	searchData := domain.SearchingTrainRequstedData{
		Date:                 req.Date,
		SourceStationid:      sourceid,
		DestinationStationid: destinationid,
	}

	res, err := h.useCase.SearchTrain(ctx, searchData)
	if err != nil {
		return nil, err
	}

	response := &pb.SearchTrainResponse{
		Traindata: make([]*pb.TrainData, len(res.SearcheResponse)), // Initialize the slice
		Status:    http.StatusOK,
	}

	for i, rs := range res.SearcheResponse {
		response.Traindata[i] = &pb.TrainData{
			Trainname: rs.TrainName,
			Time:      nil,
		}
	}

	return response, nil
}
