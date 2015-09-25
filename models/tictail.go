package models

import (
	"encoding/json"
	"errors"
	"github.com/byrnedo/tictochimp/utils"
	"time"
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

type MeStoreResponse struct {
	ID        string
	Subdomain string
}

type MeResponse struct {
	Email  string
	Stores []MeStoreResponse
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

type OrdersResponse struct {
	Customer    *OrdersCustomerResponse
	Transaction *OrdersTransactionResponse
}

type OrdersCustomerResponse struct {
	Name       string
	Language   string
	Country    string
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	ID         string
	Email      string
}
type OrdersTransactionResponse struct {
	Reference     string
	Status        string
	PaidAt        time.Time `json:"paid_at"`
	PendingReason string    `json:"pending_reason"`
	Processor     string
}

type OrderItemResponse struct {
	Currency string
	Price    int
	Product  *OrderItemProductResponse
	Quantity int
}

type OrderItemProductResponse struct {
	Title       string
	Status      string
	StoreUrl    string `json:"store_url"`
	Description string
	StoreID     string    `json:"store_id"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Price       int
	Currency    string
	ID          string
}

func (m *Tictail) GetAllOrders(storeID string) (*OrdersResponse, error) {

	responseData := OrdersResponse{}

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
