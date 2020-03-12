package repo

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"iiujapp.tech/susaday/model"
)

// ProductRepository struct
type ProductRepository struct {
	db *sql.DB
}

// NewProduct start handle
func NewProduct(db *sql.DB) ProductRepository {
	return ProductRepository{
		db: db,
	}
}

// FindProduct function
func (repo ProductRepository) FindProduct(proID string) (model.Product, error) {
	stmt, err := repo.db.Prepare(`
				SELECT 
					id,
					category_id,
					code,
					topic,
					detail,
					base_price,
					price,
					stock_type,
					stock,
					status,
					position,
					created_at,
					updated_at
				FROM product
				WHERE id = ?
			`)

	if err != nil {
		return model.Product{}, err
	}
	defer stmt.Close()

	var item model.Product
	err = stmt.QueryRow(proID).Scan(
		&item.ID,
		&item.CategoryID,
		&item.Code,
		&item.Topic,
		&item.Detail,
		&item.BasePrice,
		&item.Price,
		&item.StockType,
		&item.Stock,
		&item.Status,
		&item.Position,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	return item, err
}
