package models

import "time"

type Product struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	ProductCategoryID int       `json:"product_category_id"`
	Name              string    `json:"name"`
	Price             int       `json:"price"`
	Qty               int       `json:"qty"`
	Desc              string    `json:"desc"`
	Thumbnail         string    `json:"thumbnail"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
