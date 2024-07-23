package repository

import (
	"grpc-practice/config"
	"grpc-practice/grpc/client"
	auth "grpc-practice/grpc/proto"
)

type Repository struct {
	config *config.Config
	client *client.GRPCClient
}

func NewRepository(config *config.Config) *Repository {
	return &Repository{config: config}
}

func (repository Repository) CreateAuth(name string) (*auth.AuthData, error) {
	return repository.client.CreateAuth(name)
}
