package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	DateTime    time.Time `json:"datetime"`
	Location    string    `json:"location"`
	Price       float64   `json:"price"`
	ImgURL      string    `json:"img_url"`
	Category    string    `json:"category"`
	Priority    int       `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}