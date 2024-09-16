package models

import "time"

type InvoiceDetail struct {
	ID        int       `json:"id"`
	InvoiceID int       `json:"invoice_id"`
	ProductID int       `json:"product_id"`
	Price     int       `json:"price"`
	Qty       int       `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
