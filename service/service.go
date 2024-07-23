package service

import (
	"grpc-practice/config"
	auth "grpc-practice/grpc/proto"
	"grpc-practice/repository"
)

type Service struct {
	config     *config.Config
	repository *repository.Repository
}

func NewService(config *config.Config, repository *repository.Repository) *Service {
	return &Service{config: config, repository: repository}
}

func (service *Service) CreateAuth(name string) (*auth.AuthData, error) {
	return service.repository.CreateAuth(name)
}
