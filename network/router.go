package network

import (
	"github.com/gin-gonic/gin"
	"grpc-practice/config"
	"grpc-practice/grpc/client"
	"grpc-practice/service"
)

type Network struct {
	client  *client.GRPCClient
	config  *config.Config
	service *service.Service
	engin   *gin.Engine
}

func NewNetwork(config *config.Config, service *service.Service, client *client.GRPCClient) *Network {
	network := &Network{config: config, service: service, engin: gin.New(), client: client}

	network.engin.POST(network.login())
	network.engin.GET(network.verify())

	return network
}

func (network *Network) StartServer(address string) {
	network.engin.Run(address)
}
