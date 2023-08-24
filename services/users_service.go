package services

import (
	"main/gateways"
	"main/gateways/models"
)

type UsersService interface {
	GetUsers() (*models.UsersResponse, error)
}

type usersService struct {
	usersGateways gateways.UsersGateways
}

func InitUsersService(usersGateways gateways.UsersGateways) UsersService {
	return &usersService{
		usersGateways,
	}
}

func (s *usersService) GetUsers() (*models.UsersResponse, error) {
	users, err := s.usersGateways.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
