package models

import (
	"encoding/json"
	"errors"
	"github.com/byrnedo/tictochimp/models/tictailSpec"
	"github.com/byrnedo/tictochimp/utils"
)

type Tictail struct {
	url         string
	accessToken string
	client      *utils.RestClient
}

const TICTAIL_API_VERSION = "v1"
const TICTAIL_API_URL = "https://api.tictail.com"

func NewTictail(accessToken string) *Tictail {

	client := utils.NewRestClient()

	client.Headers = map[string]string{
		"Authorization": "Bearer " + accessToken,
	}
	return &Tictail{
		TICTAIL_API_URL + "/" + TICTAIL_API_VERSION,
		accessToken,
		client,
	}
}

func (m *Tictail) GetMe() (*tictailSpec.MeResponse, error) {

	responseData := tictailSpec.MeResponse{}

	err := m.client.Get(m.url + "/me")
	if err != nil {
		return nil, errors.New("Failed to do GET request: " + err.Error())
	}

	status := m.client.ResponseStatus()
	if status != "200" && status != "201" {
		return nil, errors.New("Got non-success status code: " + status)
	}

	err = json.Unmarshal(m.client.ResponseBody(), &responseData)

	return &responseData, nil
}

func (m *Tictail) GetAllOrders(storeID string) (*tictailSpec.OrdersResponse, error) {

	responseData := tictailSpec.OrdersResponse{}

	err := m.client.Get(m.url + "/" + storeID + "/orders")
	if err != nil {
		return nil, errors.New("Failed to do GET request: " + err.Error())
	}

	status := m.client.ResponseStatus()
	if status != "200" && status != "201" {
		return nil, errors.New("Got non-success status code: " + status)
	}

	err = json.Unmarshal(m.client.ResponseBody(), &responseData)

	return &responseData, nil
}
