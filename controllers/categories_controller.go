package controllers

import (
	
	"tacna-events-backend/services"
	"log"
	"strconv"
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
func (cc *CategoryController) CreateCategory(c *gin.Context) {
	log.Println("CreateCategory called, body: ", c.Request.Body)
	var category struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if err := cc.categoryService.CreateCategory(category.Name); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(201, gin.H{"message": "Category created successfully"})
	
}

func (cc *CategoryController) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	categoryID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid category ID"})
		return
	}
	category, err := cc.categoryService.GetCategory(uint(categoryID))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get category"})
		return
	}
	c.JSON(200, category)
	
}

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	var category struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid category ID"})
		return
	}
	if err := cc.categoryService.UpdateCategory(uint(categoryID), category.Name); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update category"})
		return
	}
	c.JSON(200, gin.H{"message": "Category updated successfully"})
	
}


