package api

import (
	"log"
	"net"

	"github.com/athunlal/bookNowTrain-svc/pkg/api/handler"
	"github.com/athunlal/bookNowTrain-svc/pkg/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServerHttp struct {
	Engine *gin.Engine
}

func NewGrpcServer(TrainHandler *handler.TrainHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalln("Failed to listen to the GRPC Port", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTrainManagementServer(grpcServer, TrainHandler)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Could not serve the GRPC Server: ", err)
	}
}

func NewServerHttp(trainHandler *handler.TrainHandler) *ServerHttp {
	engine := gin.New()

	go NewGrpcServer(trainHandler, "8892")

	engine.Use(gin.Logger())
	return &ServerHttp{
		Engine: engine,
	}
}

func (ser *ServerHttp) Start() {
	ser.Engine.Run(":9000")
}
