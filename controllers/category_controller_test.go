package controllers

import (
	"deals/services"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCategoriesController_ListCategories(t *testing.T) {
	type fields struct {
		categoryService *services.CategoryService
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
			cc := &CategoriesController{
				categoryService: tt.fields.categoryService,
			}
			cc.ListCategories(tt.args.c)
		})
	}
}

func TestCategoriesController_AddCategoryForm(t *testing.T) {
	type fields struct {
		categoryService *services.CategoryService
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
			cc := &CategoriesController{
				categoryService: tt.fields.categoryService,
			}
			cc.AddCategoryForm(tt.args.c)
		})
	}
}

func TestCategoriesController_AddCategory(t *testing.T) {
	type fields struct {
		categoryService *services.CategoryService
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
			cc := &CategoriesController{
				categoryService: tt.fields.categoryService,
			}
			cc.AddCategory(tt.args.c)
		})
	}
}

func TestCategoriesController_EditCategory(t *testing.T) {
	type fields struct {
		categoryService *services.CategoryService
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
			cc := &CategoriesController{
				categoryService: tt.fields.categoryService,
			}
			cc.EditCategory(tt.args.c)
		})
	}
}

func TestCategoriesController_EditCategoryForm(t *testing.T) {
	type fields struct {
		categoryService *services.CategoryService
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
			cc := &CategoriesController{
				categoryService: tt.fields.categoryService,
			}
			cc.EditCategoryForm(tt.args.c)
		})
	}
}

func TestCategoriesController_DeleteCategory(t *testing.T) {
	type fields struct {
		categoryService *services.CategoryService
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
			cc := &CategoriesController{
				categoryService: tt.fields.categoryService,
			}
			cc.DeleteCategory(tt.args.c)
		})
	}
}
