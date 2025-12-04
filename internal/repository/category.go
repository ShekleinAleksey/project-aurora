package repository

import (
	"github.com/ShekleinAleksey/project-aurora/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
}

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category entity.Category) error {
	_, err := r.db.Exec("INSERT INTO categories (name) VALUES ($1)", category.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) GetAllCategories() {

}
