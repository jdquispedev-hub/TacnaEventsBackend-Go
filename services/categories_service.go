package services

import (
	"tacna-events-backend/models"

	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (cs *CategoryService) GetCategories() ([]models.Categorie, error) {
	//un log de que entro aqui 
	//comando para ver log desde la terminal 
	//go run main.go
	var categories []models.Categorie
	result := cs.db.Find(&categories)
	return categories, result.Error
}


