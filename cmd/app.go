package cmd

import (
	"grpc-practice/config"
	"grpc-practice/grpc/client"
	"grpc-practice/network"
	"grpc-practice/repository"
	"grpc-practice/service"
)

type App struct {
	config *config.Config

	client     *client.GRPCClient
	service    *service.Service
	repository *repository.Repository
	network    *network.Network
}

func NewApp(config *config.Config) {
	app := &App{
		config: config,
	}
	app.repository = repository.NewRepository(config)
	app.service = service.NewService(config, app.repository)

	app.client, _ = client.NewGRPCClient(config)
	app.network = network.NewNetwork(config, app.service, app.client)
	app.network.StartServer("localhost:8080")
}
