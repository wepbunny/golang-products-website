package services

import (
	"database/sql"
	"deals/db"
	"deals/models"
	"fmt"
	"os"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (ps *ProductService) GetProductList() ([]models.Product, error) {

	rows, err := db.DB.Query("SELECT id, category_id, title, price, image, status, created_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.CategoryID, &product.Title, &product.Price, &product.Image, &product.Status, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (ps *ProductService) GetProductByID(id int) (*models.Product, error) {
	row := db.DB.QueryRow("SELECT * FROM products WHERE id = ?", id)

	var product models.Product
	err := row.Scan(&product.ID, &product.CategoryID, &product.Title, &product.Price, &product.Image, &product.Status, &product.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &product, nil
}

func (ps *ProductService) CreateProduct(product *models.Product) (sql.Result, error) {
	query := "INSERT INTO products (title, price, image, status) VALUES (?, ?, ?, ?)"
	result, err := db.DB.Exec(query, product.Title, product.Price, product.Image, product.Status)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ps *ProductService) UpdateProductByID(id int, product *models.Product) error {
	query := "UPDATE products SET title = ?, price = ?, category_id = ?, image = ?, status = ? WHERE id = ?"
	_, err := db.DB.Exec(query, product.Title, product.Price, product.CategoryID, product.Image, product.Status, id)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) DeleteProductByID(id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) GenerateReport() error {
	rows, err := db.DB.Query("SELECT * FROM products")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Create a new Excel file
	fileName := "products_report.csv"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the report headers
	header := []interface{}{"ID", "Category ID", "Title", "Price", "Image", "Status", "Created At"}
	err = writeRow(file, header)
	if err != nil {
		return err
	}

	// Write the product data
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.CategoryID, &product.Title, &product.Price, &product.Image, &product.Status, &product.CreatedAt)
		if err != nil {
			return err
		}

		rowData := []interface{}{product.ID, product.CategoryID, product.Title, product.Price, product.Image, product.Status, product.CreatedAt}
		err = writeRow(file, rowData)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeRow(file *os.File, rowData []interface{}) error {
	row := ""
	for _, value := range rowData {
		strValue := fmt.Sprintf("%v", value)
		row += fmt.Sprintf("%s, ", strValue)
	}
	row += "\n"

	_, err := file.WriteString(row)
	if err != nil {
		return err
	}

	return nil
}
