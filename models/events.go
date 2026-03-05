package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	DateTime    *time.Time     `json:"datetime"` // Pointer para permitir NULL
	Location    string         `json:"location"`
	Price       float64        `json:"price"`
	Image_URL   string         `json:"image_url" gorm:"column:image_url"`
	Category    string         `json:"category"`
	CategoryID  int            `json:"category_id"`
	Priority    int            `json:"priority" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete
}
