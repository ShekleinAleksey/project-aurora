package repository

import (
	"github.com/ShekleinAleksey/project-aurora/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category entity.Category) error {
	_, err := r.db.Exec("INSERT INTO categories (name) VALUES ($1)", category.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) GetAllCategories() {

}
