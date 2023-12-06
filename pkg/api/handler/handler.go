package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	"github.com/athunlal/bookNowTrain-svc/pkg/pb"
	"github.com/athunlal/bookNowTrain-svc/pkg/response"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/usecase/interface"
	"github.com/athunlal/bookNowTrain-svc/pkg/utils"
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

func (h *TrainHandler) ViewCompartment(ctx context.Context, req *pb.ViewCompartmentRequest) (*pb.ViewCompartmentResponse, error) {
	res, err := h.useCase.ViewCompartment(ctx)
	if err != nil {
		return nil, err
	}
	return mapViewCompartmentToPbResponse(res), nil
}

func mapViewCompartmentToPbResponse(res []domain.CompartmentDetails) *pb.ViewCompartmentResponse {
	var Compartments []*pb.CompartmentDetails
	for _, val := range res {
		compartment := pb.CompartmentDetails{
			CompartmentId:   val.CompartmentId.Hex(),
			CompartmentType: val.CompartmentType,
			Price:           int64(val.Price),
		}
		Compartments = append(Compartments, &compartment)
	}
	return &pb.ViewCompartmentResponse{
		Compartment: Compartments,
	}
}

func (h *TrainHandler) ViewRoute(ctx context.Context, req *pb.ViewRoutesRequest) (*pb.ViewRoutesResponse, error) {
	res, err := h.useCase.ViewRoute(ctx)
	if err != nil {
		return nil, err
	}

	var response []*pb.RouteDetails

	for _, val := range res {
		routeDetail := &pb.RouteDetails{
			RouteName: val.RouteName,
			RouteId:   val.RouteId.Hex(),
		}
		response = append(response, routeDetail)
	}

	return &pb.ViewRoutesResponse{
		RouteDetails: response,
	}, nil
}

//View all station
func (h *TrainHandler) ViewStation(ctx context.Context, req *pb.ViewRequest) (*pb.ViewStationResponse, error) {
	res, err := h.useCase.ViewStation(ctx)
	if err != nil {
		return response.HandleError(err)
	}

	stations := response.ConvertToPBStations(res.Station)

	stationResponse := &pb.ViewStationResponse{
		Status:   http.StatusOK,
		Stations: stations,
	}

	return stationResponse, nil
}

//Updating seat into the Train collection
func (h *TrainHandler) UpdateSeatIntoTrain(ctx context.Context, req *pb.UpdateSeatIntoTrainRequest) (*pb.UpdateSeatIntoTrainResponse, error) {
	updateData, err := response.MapToUpdateData(req)
	if err != nil {
		return &pb.UpdateSeatIntoTrainResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error in input validation",
		}, err
	}

	err = h.useCase.UpadateSeatInotTrain(ctx, updateData)
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

//Add New Train
func (h *TrainHandler) AddTrain(ctx context.Context, req *pb.AddTrainRequest) (*pb.AddTrainResponse, error) {
	train := response.MapToAddTrain(req)

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

//Search Train By Source Station Id And Destination Station Id
func (h *TrainHandler) SearchTrain(ctx context.Context, req *pb.SearchTrainRequest) (*pb.SearchTrainResponse, error) {
	searchData, err := utils.PrepareSearchData(req)
	if err != nil {
		return nil, err
	}
	res, err := h.useCase.SearchTrain(ctx, searchData)
	if err != nil {
		return nil, err
	}
	response := response.MapToSearchTrainResponse(res)
	return response, nil
}

//Search Train By Train Name
func (h *TrainHandler) SearchTrainByName(ctx context.Context, req *pb.SearchTrainByNameRequest) (*pb.SearchTrainByNameResponse, error) {
	res, err := h.useCase.SearchTrainByName(ctx, req.TrainName)
	if err != nil {
		return nil, err
	}
	return response.MapToSearchTrainResponseByName(res), nil
}
