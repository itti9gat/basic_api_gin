package repo

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"iiujapp.tech/susaday/model"
)

// CategoryRepository struct
type CategoryRepository struct {
	db *sql.DB
}

// NewCategory start handle
func NewCategory(db *sql.DB) CategoryRepository {
	return CategoryRepository{
		db: db,
	}
}

// FindCategory function
func (repo CategoryRepository) FindCategory(cateID string) (model.Category, error) {
	stmt, err := repo.db.Prepare(`
				SELECT 
					id,
					topic,
					level,
					level_id,
					status,
					position,
					created_at,
					updated_at
				FROM category
				WHERE id = ?
			`)
	if err != nil {
		return model.Category{}, err
	}
	defer stmt.Close()

	var item model.Category
	err = stmt.QueryRow(cateID).Scan(
		&item.ID,
		&item.Topic,
		&item.Level,
		&item.LevelID,
		&item.Status,
		&item.Position,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	return item, err
}
