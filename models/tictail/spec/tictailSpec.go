package spec

import (
	"time"
)

type TictailTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05.999999"

func (ct *TictailTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}

	bString := string(b)
	if bString == "null" {
		return
	}
	ct.Time, err = time.Parse(ctLayout, bString)
	return
}

func (ct *TictailTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(ctLayout)), nil
}

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
	CreatedAt  TictailTime `json:"created_at"`
	ModifiedAt TictailTime `json:"modified_at"`
	ID         string
	Email      string
}
type OrdersTransactionResponse struct {
	Reference     string
	Status        string
	PaidAt        TictailTime `json:"paid_at"`
	PendingReason string      `json:"pending_reason"`
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
	StoreID     string      `json:"store_id"`
	CreatedAt   TictailTime `json:"created_at"`
	ModifiedAt  TictailTime `json:"modified_at"`
	Price       int
	Currency    string
	ID          string
}
