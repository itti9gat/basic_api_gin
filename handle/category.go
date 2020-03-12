package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CategoryByID Handle
func (h *Handler) CategoryByID(c *gin.Context) {
	cateid := c.Param("cateid")
	v, _ := h.CategoryRepository.FindCategory(cateid)
	c.JSON(http.StatusOK, v)
}
