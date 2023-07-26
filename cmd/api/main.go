package main

import (
	"log"

	"github.com/athunlal/bookNowTrain-svc/pkg/config"
	"github.com/athunlal/bookNowTrain-svc/pkg/di"
)

func main() {
	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		log.Fatal("Could not load the config file :", cfgErr)
		return
	}
	server, err := di.InitApi(cfg)
	if err != nil {
		log.Fatalln("Error in initializing the API", err)
	}
	server.Start()
}
