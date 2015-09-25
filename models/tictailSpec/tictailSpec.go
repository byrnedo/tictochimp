package tictailSpec

import (
	"time"
)

type MeStoreResponse struct {
	ID        string
	Subdomain string
}

type MeResponse struct {
	Email  string
	Stores []MeStoreResponse
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
