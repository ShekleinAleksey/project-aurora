package repository

import (
	"fmt"
	"time"

	"github.com/ShekleinAleksey/project-aurora/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	Create(category entity.CreateCategoryRequest) (int, error)
	GetAll() ([]entity.Category, error)
	GetByID(id int) (entity.Category, error)
	Update(category entity.Category) error
	Delete(id int) error
}

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category entity.CreateCategoryRequest) (int, error) {
	var id int
	now := time.Now()
	err := r.db.QueryRow("INSERT INTO categories (name, createdat) VALUES ($1, $2) RETURNING id", category.Name, now).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *categoryRepository) GetAll() ([]entity.Category, error) {
	var categories []entity.Category
	query := "SELECT * FROM categories"
	err := r.db.Select(&categories, query)
	if err != nil {
		return []entity.Category{}, fmt.Errorf("failed to get all categories: %w", err)
	}
	return categories, nil
}

func (r *categoryRepository) GetByID(id int) (entity.Category, error) {
	var category entity.Category
	query := "SELECT * FROM categories WHERE id=$1"
	err := r.db.Get(&category, query, id)
	if err != nil {
		return entity.Category{}, fmt.Errorf("failed to get category by id: %w", err)
	}

	return category, nil
}

func (r *categoryRepository) Update(category entity.Category) error {
	query := "UPDATE categories SET name=$1 WHERE id=$2"
	_, err := r.db.Exec(query, category.Name, category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id=$1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
