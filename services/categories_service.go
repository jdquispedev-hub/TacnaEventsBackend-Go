package services

import (
	"tacna-events-backend/models"
	"log"

	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (cs *CategoryService) GetCategories() ([]models.Category, error) {
	//un log de que entro aqui 
	//comando para ver log desde la terminal 
	//go run main.go
	// log.Println("GetCategories called")
	log.Print("entroooo")

	var categories []models.Category
	result := cs.db.Find(&categories)
	log.Print("categories: ", categories)
	log.Print("result.Error: ", result.Error)
	return categories, result.Error
}

func (cs *CategoryService) CreateCategory(name string) error {
	category := models.Category{Name: name}
	result := cs.db.Create(&category)
	log.Print("category: ", category)
	log.Print("result.Error: ", result.Error)
	return result.Error
}

func (cs *CategoryService) GetCategory(id uint) (models.Category, error) {
	var category models.Category
	result := cs.db.First(&category, id)
	log.Print("category: ", category)
	log.Print("result.Error: ", result.Error)
	return category, result.Error
}

func (cs *CategoryService) UpdateCategory(id uint, name string) error {
	category := models.Category{Name: name}
	result := cs.db.Model(&category).Where("id = ?", id).Update("name", name)
	log.Print("category: ", category)
	log.Print("result.Error: ", result.Error)
	return result.Error
}



