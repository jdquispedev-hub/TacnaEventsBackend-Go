package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateTime    time.Time `json:"datetime"`
	Location    string    `json:"location"`
	Price       float64   `json:"price"`
	Image_URL   string    `json:"image_url"`
	Category    string    `json:"category"`
	Priority    int       `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}