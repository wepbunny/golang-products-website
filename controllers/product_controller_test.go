package controllers

import (
	"deals/services"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestProductsController_ListProducts(t *testing.T) {
	type fields struct {
		productService *services.ProductService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProductsController{
				productService: tt.fields.productService,
			}
			pc.ListProducts(tt.args.c)
		})
	}
}

func TestProductsController_GetProduct(t *testing.T) {
	type fields struct {
		productService *services.ProductService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProductsController{
				productService: tt.fields.productService,
			}
			pc.GetProduct(tt.args.c)
		})
	}
}

func TestProductsController_ListAdminProducts(t *testing.T) {
	type fields struct {
		productService *services.ProductService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProductsController{
				productService: tt.fields.productService,
			}
			pc.ListAdminProducts(tt.args.c)
		})
	}
}

func TestProductsController_GetAdminProduct(t *testing.T) {
	type fields struct {
		productService *services.ProductService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProductsController{
				productService: tt.fields.productService,
			}
			pc.GetAdminProduct(tt.args.c)
		})
	}
}

func TestProductsController_AddProductForm(t *testing.T) {
	type fields struct {
		productService *services.ProductService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProductsController{
				productService: tt.fields.productService,
			}
			pc.AddProductForm(tt.args.c)
		})
	}
}

func TestProductsController_AddProduct(t *testing.T) {
	type fields struct {
		productService *services.ProductService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProductsController{
				productService: tt.fields.productService,
			}
			pc.AddProduct(tt.args.c)
		})
	}
}

func TestProductsController_EditProductForm(t *testing.T) {
	type fields struct {
		productService *services.ProductService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProductsController{
				productService: tt.fields.productService,
			}
			pc.EditProductForm(tt.args.c)
		})
	}
}

func TestProductsController_EditProduct(t *testing.T) {
	type fields struct {
		productService *services.ProductService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProductsController{
				productService: tt.fields.productService,
			}
			pc.EditProduct(tt.args.c)
		})
	}
}

func TestProductsController_DeleteProduct(t *testing.T) {
	type fields struct {
		productService *services.ProductService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProductsController{
				productService: tt.fields.productService,
			}
			pc.DeleteProduct(tt.args.c)
		})
	}
}

func TestProductsController_GenerateReport(t *testing.T) {
	type fields struct {
		productService *services.ProductService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &ProductsController{
				productService: tt.fields.productService,
			}
			pc.GenerateReport(tt.args.c)
		})
	}
}
