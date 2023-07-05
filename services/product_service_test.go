package services

import (
	"database/sql"
	"deals/models"
	"os"
	"reflect"
	"testing"
)

func TestProductService_GetProductList(t *testing.T) {
	tests := []struct {
		name    string
		ps      *ProductService
		want    []models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ProductService{}
			got, err := ps.GetProductList()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetProductList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.GetProductList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_GetProductByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		ps      *ProductService
		args    args
		want    *models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ProductService{}
			got, err := ps.GetProductByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.GetProductByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_CreateProduct(t *testing.T) {
	type args struct {
		product *models.Product
	}
	tests := []struct {
		name    string
		ps      *ProductService
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ProductService{}
			got, err := ps.CreateProduct(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_UpdateProductByID(t *testing.T) {
	type args struct {
		id      int
		product *models.Product
	}
	tests := []struct {
		name    string
		ps      *ProductService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ProductService{}
			if err := ps.UpdateProductByID(tt.args.id, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("ProductService.UpdateProductByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductService_DeleteProductByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		ps      *ProductService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ProductService{}
			if err := ps.DeleteProductByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ProductService.DeleteProductByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductService_GenerateReport(t *testing.T) {
	tests := []struct {
		name    string
		ps      *ProductService
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ProductService{}
			if err := ps.GenerateReport(); (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GenerateReport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_writeRow(t *testing.T) {
	type args struct {
		file    *os.File
		rowData []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeRow(tt.args.file, tt.args.rowData); (err != nil) != tt.wantErr {
				t.Errorf("writeRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
