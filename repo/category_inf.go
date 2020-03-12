//go:generate mockgen -destination ./mock/category_inf_mock.go iiujapp.tech/susaday/repo CategoryINF

package repo

import (
	"iiujapp.tech/susaday/model"
)

// CategoryINF interface
type CategoryINF interface {
	FindCategory(proID string) (model.Category, error)
}
