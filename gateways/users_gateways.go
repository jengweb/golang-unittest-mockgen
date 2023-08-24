package gateways

import (
	"encoding/json"
	"io"
	"main/gateways/models"
	"net/http"
)

type UsersGateways interface {
	GetUsers() (*models.UsersResponse, error)
}

type usersGateways struct {
}

func InitUsersGateways() UsersGateways {
	return &usersGateways{}
}

func (g usersGateways) GetUsers() (*models.UsersResponse, error) {
	req, err := http.NewRequest("GET", "https://reqres.in/api/users", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var users models.UsersResponse
	json.Unmarshal(data, &users)
	return &users, nil
}
