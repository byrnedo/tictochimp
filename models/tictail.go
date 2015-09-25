package models

import (
	"encoding/json"
	"errors"
	"github.com/byrnedo/tictochimp/utils"
)

type Tictail struct {
	url         string
	accessToken string
	client      *utils.RestClient
}

const TICTAIL_API_VERSION = "1.0"

func NewTictail(url string, accessToken string) *Tictail {

	client := utils.NewRestClient()

	client.Headers = map[string]string{
		"Authorization": "Bearer " + accessToken,
	}
	return &Tictail{
		url + "/" + API_VERSION,
		accessToken,
		client,
	}
}

type MeStoreResponse struct {
	ID        string
	Subdomain string
}

type MeResponse struct {
	Email  string
	Stores string
}

func (m *Tictail) GetMe() (*MeResponse, error) {

	responseData := MeResponse{}

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
