//go:generate mockgen -destination ./mock/product_inf_mock.go iiujapp.tech/susaday/repo ProductINF

package repo

import (
	"iiujapp.tech/basic-api-gin/model"
)

// ProductINF interface
type ProductINF interface {
	FindProduct(proID string) (model.Product, error)
}
