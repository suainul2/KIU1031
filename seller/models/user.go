package models

import "time"

type User struct {
	ID              int       `json:"id"`
	SubdistrictID   int       `json:"subdistrict_id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Role            int       `json:"role"`
	Phone           string    `json:"phone"`
	Address         string    `json:"address"`
	Image           string    `json:"image"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

const USER_ROLE_CUSTOMER = 1
const USER_ROLE_SELLER = 2
const USER_ROLE_DRIVER = 3
