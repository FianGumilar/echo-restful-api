package repository

import (
	"context"
	"database/sql"

	"github.com/FianGumilar/restful-api-echo/interfaces"
	"github.com/FianGumilar/restful-api-echo/models/dto"
)

type repository struct {
	db *sql.DB
}

func NewCategoryRepository(con *sql.DB) interfaces.CategoryRepository {
	return &repository{db: con}
}

// Create implements interfaces.CategoryRepository.
func (r repository) Create(ctx context.Context, category *dto.Category) dto.Category {
	query := `INSERT INTO categories (name) VALUES (?) RETURNING id`

	res, err := r.db.ExecContext(ctx, query, category.Name)
	if err != nil {
		panic(err)
	}

	lastInsertID, _ := res.LastInsertId()
	id := int64(lastInsertID)

	category.SetID(&id)

	return *category
}

// Delete implements interfaces.CategoryRepository.
func (r repository) Delete(ctx context.Context, category *dto.Category) error {
	query := `DELETE FROM categories WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, category.GetID())
	if err != nil {
		panic(err.Error())
	}
	return nil
}

// FindAll implements interfaces.CategoryRepository.
func (r repository) FindAll(ctx context.Context) ([]dto.Category, error) {
	query := `SELECT id, name FROM categories`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		panic(err.Error())
	}

	var categories []dto.Category

	for rows.Next() {
		var category dto.Category
		rows.Scan(category.GetID(), category.GetName())
		categories = append(categories, category)
	}
	return categories, nil
}

// FindByID implements interfaces.CategoryRepository.
func (r repository) FindByID(ctx context.Context, id int64) (dto.Category, error) {
	query := `SELECT id, name FROM categories WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	var category dto.Category

	err := row.Scan(category.GetID(), category.GetName())
	if err != nil {
		panic(err.Error())
	}

	return category, nil
}

// Update implements interfaces.CategoryRepository.
func (r repository) Update(ctx context.Context, category *dto.Category) error {
	query := `UPDATE catgories SET name = ? WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, category.GetName(), category.GetID())

	if err != nil {
		panic(err.Error())
	}
	return nil
}
