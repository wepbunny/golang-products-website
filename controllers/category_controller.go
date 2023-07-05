package controllers

import (
	"deals/models"
	"deals/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	categoryService *services.CategoryService
}

func NewCategoriesController() *CategoriesController {
	return &CategoriesController{}
}

func (cc *CategoriesController) ListCategories(c *gin.Context) {

	categories, err := cc.categoryService.GetCategoryList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to load: %v", err)})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (cc *CategoriesController) AddCategoryForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add_category.html", nil)
}

func (cc *CategoriesController) AddCategory(c *gin.Context) {
	var category models.Category
	// Parse form data
	category.Name = c.PostForm("name")

	result, err := cc.categoryService.CreateCategory(&category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add category: %v", err)})
		return
	}

	// Get the auto-generated ID of the new category
	categoryID, err := result.LastInsertId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrive category id: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Category added successfully with ID %d", categoryID)})
}

func (cc *CategoriesController) EditCategoryForm(c *gin.Context) {
	id := c.Param("id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category ID"})
		return
	}

	category, err := cc.categoryService.GetCategoryByID(categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Category not found: %v", err)})
		return
	}

	c.HTML(http.StatusOK, "edit_category.html", category)
}

func (cc *CategoriesController) EditCategory(c *gin.Context) {
	id := c.Param("id")
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	var category models.Category

	// Parse form data
	category.Name = c.PostForm("name")

	err = cc.categoryService.UpdateCategoryByID(categoryID, &category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func (cc *CategoriesController) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	err = cc.categoryService.DeleteCategoryByID(categoryID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
