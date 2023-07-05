package services

import (
	"database/sql"
	"deals/db"
	"deals/models"
)

type CategoryService struct{}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (cs *CategoryService) GetCategoryList() ([]models.Category, error) {
	var categories []models.Category

	query := "SELECT * FROM categories"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (cs *CategoryService) GetCategoryByID(id int) (*models.Category, error) {
	var category models.Category

	query := "SELECT * FROM categories WHERE id = ?"
	err := db.DB.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (cs *CategoryService) CreateCategory(category *models.Category) (sql.Result, error) {
	query := "INSERT INTO categories (name) VALUES (?)"
	result, err := db.DB.Exec(query, category.Name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cs *CategoryService) UpdateCategoryByID(id int, category *models.Category) error {
	query := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := db.DB.Exec(query, category.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CategoryService) DeleteCategoryByID(id int) error {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
