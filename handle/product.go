package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProductByID Handle
func (h *Handler) ProductByID(c *gin.Context) {
	proid := c.Param("proid")
	v, _ := h.ProductRepository.FindProduct(proid)
	c.JSON(http.StatusOK, v)
}
