package controllers

import (
	"deals/models"
	"deals/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	productService *services.ProductService
}

func NewProductsController() *ProductsController {
	return &ProductsController{}
}

func (pc *ProductsController) ListProducts(c *gin.Context) {
	products, err := pc.productService.GetProductList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to load: %v", err)})
		return
	}

	c.HTML(http.StatusOK, "products.html", products)
}

func (pc *ProductsController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := pc.productService.GetProductByID(productID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.HTML(http.StatusOK, "product.html", product)
}

func (pc *ProductsController) ListAdminProducts(c *gin.Context) {
	products, err := pc.productService.GetProductList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to load: %v", err)})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductsController) GetAdminProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := pc.productService.GetProductByID(productID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (pc *ProductsController) AddProductForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add_product.html", nil)
}

func (pc *ProductsController) AddProduct(c *gin.Context) {
	var product models.Product
	priceStr := c.PostForm("price")
	categoryIDStr := c.PostForm("category_id")
	// Convert price and category ID to appropriate types
	price, _ := strconv.ParseFloat(priceStr, 64)
	categoryID, _ := strconv.Atoi(categoryIDStr)
	// Parse form data
	product.Title = c.PostForm("title")
	product.Price = price
	product.CategoryID = categoryID
	product.Image = c.PostForm("image")
	product.Status = c.PostForm("status")
	result, err := pc.productService.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add product: %v", err)})
		return
	}

	// Get the auto-generated ID of the new product
	productID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Product added successfully with ID %d", productID)})
}

func (pc *ProductsController) EditProductForm(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := pc.productService.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Product not found: %v", err)})
		return
	}

	c.HTML(http.StatusOK, "edit_product.html", product)
}

func (pc *ProductsController) EditProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	priceStr := c.PostForm("price")
	categoryIDStr := c.PostForm("category_id")
	// Convert price and category ID to appropriate types
	price, _ := strconv.ParseFloat(priceStr, 64)
	categoryID, _ := strconv.Atoi(categoryIDStr)
	// Parse form data
	product.Title = c.PostForm("title")
	product.Price = price
	product.CategoryID = categoryID
	product.Image = c.PostForm("image")
	product.Status = c.PostForm("status")
	err = pc.productService.UpdateProductByID(productID, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update product: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (pc *ProductsController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	err = pc.productService.DeleteProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete product: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (pc *ProductsController) GenerateReport(c *gin.Context) {
	err := pc.productService.GenerateReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to generate report: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Report generated successfully"})
}
