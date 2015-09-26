package tictail

import (
	"encoding/json"
	"errors"
	"github.com/byrnedo/tictochimp/models/tictail/spec"
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

func (m *Tictail) GetMe() (*spec.MeResponse, error) {

	responseData := spec.MeResponse{}

	err := m.client.Get(m.url + "/me")
	if err != nil {
		return nil, errors.New("Failed to do GET request: " + err.Error())
	}

	status := m.client.ResponseStatus()
	if status != 200 && status != 201 {
		return nil, errors.New("Got non-success status code: " + string(status))
	}

	err = json.Unmarshal(m.client.ResponseBody(), &responseData)

	return &responseData, err
}

func (m *Tictail) GetAllOrders(storeID string) ([]spec.OrdersResponse, error) {

	responseData := []spec.OrdersResponse{}

	err := m.client.Get(m.url + "/stores/" + storeID + "/orders")
	if err != nil {
		return nil, errors.New("Failed to do GET request: " + err.Error())
	}

	status := m.client.ResponseStatus()
	if status != 200 && status != 201 {
		return nil, errors.New("Got non-success status code: " + string(status))
	}

	err = json.Unmarshal(m.client.ResponseBody(), &responseData)

	return responseData, err
}
