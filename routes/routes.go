package routes

import (
	"deals/controllers"

	"github.com/gin-gonic/gin"
)

var (
	productsController   = controllers.NewProductsController()
	categoriesController = controllers.NewCategoriesController()
)

func InitRoutes(r *gin.Engine) {
	r.Static("/public", "./public")
	r.LoadHTMLGlob("views/*")
	r.GET("/products", productsController.ListProducts)
	r.GET("/products/:id", productsController.GetProduct)
	r.GET("/admin/products", productsController.ListAdminProducts)
	r.GET("/admin/products/:id", productsController.GetAdminProduct)
	r.GET("/admin/products/add", productsController.AddProductForm)
	r.POST("/admin/products/add", productsController.AddProduct)
	r.GET("/admin/products/edit/:id", productsController.EditProductForm)
	r.POST("/admin/products/edit/:id", productsController.EditProduct)
	r.GET("/admin/products/delete/:id", productsController.DeleteProduct)
	r.GET("/admin/categories", categoriesController.ListCategories)
	r.GET("/admin/categories/add", categoriesController.AddCategoryForm)
	r.POST("/admin/categories/add", categoriesController.AddCategory)
	r.GET("/admin/categories/edit/:id", categoriesController.EditCategoryForm)
	r.POST("/admin/categories/edit/:id", categoriesController.EditCategory)
	r.GET("/admin/categories/delete/:id", categoriesController.DeleteCategory)
	r.GET("/report", productsController.GenerateReport)
}
