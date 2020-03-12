package route

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"iiujapp.tech/susaday/handle"
)

// NewRoute main route
func NewRoute(r *gin.Engine, h *handle.Handler) {
	r.GET("/v1/category/:cateid", h.CategoryByID)
	r.GET("/v1/product/:proid", h.ProductByID)
}
