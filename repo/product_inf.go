//go:generate mockgen -destination ./mock/product_inf_mock.go iiujapp.tech/susaday/repo ProductINF

package repo

import (
	"iiujapp.tech/susaday/model"
)

// ProductINF interface
type ProductINF interface {
	FindProduct(proID string) (model.Product, error)
}
