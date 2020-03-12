package model

// Product struct
type Product struct {
	ID         int64   `json:"id"`
	CategoryID int64   `json:"category_id"`
	Code       string  `json:"code"`
	Topic      string  `json:"topic"`
	Detail     string  `json:"detail"`
	BasePrice  float64 `json:"base_price"`
	Price      float64 `json:"price"`
	StockType  string  `json:"stock_type"`
	Stock      int64   `json:"stock"`
	Status     string  `json:"status"`
	Position   int64   `json:"position"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
