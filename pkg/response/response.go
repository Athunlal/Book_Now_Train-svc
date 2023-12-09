package response

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/athunlal/bookNowTrain-svc/pkg/domain"
	"github.com/athunlal/bookNowTrain-svc/pkg/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapToSearchTrainResponseByName(res domain.Train) *pb.SearchTrainByNameResponse {
	return &pb.SearchTrainByNameResponse{
		TrainName:    res.TrainName,
		TrainType:    res.TrainType,
		TranNumber:   strconv.FormatInt(int64(res.TrainNumber), 10),
		Startingtime: res.StartingTime,
		Endingtime:   res.EndingtingTime,
	}
}

func MapToSearchTrainResponse(res domain.SearchingTrainResponseData) *pb.SearchTrainResponse {
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

	return response
}

func MapToAddTrain(req *pb.AddTrainRequest) domain.Train {
	date := make([]domain.Date, len(req.Date))
	for i, val := range req.Date {
		date[i].Day = val.Day
	}
	return domain.Train{
		TrainName:      req.Trainname,
		TrainNumber:    uint(req.Trainnumber),
		TrainType:      req.Traintype,
		StartingTime:   req.Startingtime,
		EndingtingTime: req.Endingtime,
		Date:           date,
	}
}

func MapToUpdateData(req *pb.UpdateSeatIntoTrainRequest) (domain.Train, error) {
	updateData := domain.Train{
		TrainNumber: uint(req.Trainnumber),
		Compartment: make([]domain.Compartment, len(req.Compartments)),
	}

	for i, rs := range req.Compartments {
		seatid, err := primitive.ObjectIDFromHex(rs.Seatid)
		if err != nil {
			return domain.Train{}, fmt.Errorf("converting string to primitive.ObjectID: %v", err)
		}
		updateData.Compartment[i] = domain.Compartment{
			Seatid: seatid,
		}
	}

	return updateData, nil
}

func HandleError(err error) (*pb.ViewStationResponse, error) {
	return &pb.ViewStationResponse{
		Status: http.StatusUnprocessableEntity,
		Error:  "Error Found in usecase",
	}, err
}

func ConvertToPBStations(domainStations []domain.Station) []*pb.Station {
	stations := make([]*pb.Station, len(domainStations))

	for i, station := range domainStations {
		pbStation := &pb.Station{
			Stationid:   station.StationId.Hex(),
			StationName: station.StationName,
			City:        station.City,
		}
		stations[i] = pbStation
	}

	return stations
}
