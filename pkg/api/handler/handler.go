package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	"github.com/athunlal/bookNowTrain-svc/pkg/pb"
	interfaces "github.com/athunlal/bookNowTrain-svc/pkg/usecase/interface"
)

type TrainHandler struct {
	useCase interfaces.TrainUseCase
	pb.TrainManagementServer
}

type StationHandler struct {
	useCase interfaces.SationUseCase
	pb.TrainManagementServer
}

func NewTrainHandler(usecase interfaces.TrainUseCase) *TrainHandler {
	return &TrainHandler{
		useCase: usecase,
	}
}

func (h *StationHandler) AddStation(ctx context.Context, req *pb.AddStationRequest) (*pb.AddStationResponse, error) {
	fmt.Println("this is station handler----------->>>>>>>>>>")
	station := domain.Station{
		StationId:   uint(req.Stationid),
		StationName: req.Stationname,
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
