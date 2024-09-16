package models

import "time"

type Invoice struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	DriverID      int       `json:"driver_id"`
	Code          string    `json:"code"`
	ShippingCost  int       `json:"shipping_cost"`
	FromLatitude  float64   `json:"from_latitude"`
	FromLongitude float64   `json:"from_longitude"`
	ToLatitude    float64   `json:"to_latitude"`
	ToLongitude   float64   `json:"to_longitude"`
	Address       string    `json:"address"`
	Status        int       `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

const INVOICE_STATUS_SEARCH_DRIVER = 1
const INVOICE_STATUS_PROCESS = 2
const INVOICE_STATUS_DELIVERY = 3
const INVOICE_STATUS_FINISH = 4
const INVOICE_STATUS_CANCEL = 5
