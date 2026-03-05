package controllers

import (
	"tacna-events-backend/services"
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{
		categoryService: services.NewCategoryService(db),
	}
}
func (cc *CategoryController) GetCategories(c *gin.Context) {
	log.Println("GetCategories called")
	categories, err := cc.categoryService.GetCategories()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get categories"})
		return
	}
	c.JSON(200, categories)
}
