//go:build wireinject
// +build wireinject

package di

import (
	"github.com/athunlal/bookNowTrain-svc/pkg/api"
	"github.com/athunlal/bookNowTrain-svc/pkg/api/handler"
	"github.com/athunlal/bookNowTrain-svc/pkg/config"
	"github.com/athunlal/bookNowTrain-svc/pkg/db"
	"github.com/athunlal/bookNowTrain-svc/pkg/repository"
	"github.com/athunlal/bookNowTrain-svc/pkg/usecase"
	"github.com/google/wire"
)

func InitApi(cfg config.Config) (*api.ServerHttp, error) {
	wire.Build(
		db.ConnectDataBase,
		repository.NewTrainRepo,
		usecase.NewTrainUseCase,
		handler.NewTrainHandler,
		api.NewServerHttp)
	return &api.ServerHttp{}, nil
}
