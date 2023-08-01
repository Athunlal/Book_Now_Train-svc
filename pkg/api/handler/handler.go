package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	"github.com/athunlal/bookNowTrain-svc/pkg/pb"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/usecase/interface"
	"github.com/golang/protobuf/ptypes/timestamp"
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
	routeData := domain.Route{}
	for _, rs := range req.Station {
		stationID := primitive.NewObjectIDFromTimestamp(time.Unix(rs.Time.Seconds, int64(rs.Time.Nanos)))
		distance := float64(rs.Distance)

		// Convert to a time.Time value

		routeData.RouteMap = append(routeData.RouteMap, domain.RouteStation{
			StationId: stationID,
			Distance:  distance,
			Time: &timestamp.Timestamp{
				Seconds: rs.Time.Seconds,
				Nanos:   int32(rs.Time.Nanos),
			},
		})
	}
	err := h.useCase.AddRoute(ctx, routeData)
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
		StationId:   uint(req.Stationid),
		StationName: req.Stationname,
		City:        req.City,
	}

	err := h.useCase.AddStation(ctx, station)
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
