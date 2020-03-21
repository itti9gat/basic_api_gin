package handle

import "iiujapp.tech/basic-api-gin/repo"

// Handler struct
type Handler struct {
	CategoryRepository repo.CategoryINF
	ProductRepository  repo.ProductINF
}

// NewHandler struct
func NewHandler(rb repo.CategoryINF, rp repo.ProductINF) *Handler {
	return &Handler{
		CategoryRepository: rb,
		ProductRepository:  rp,
	}
}
