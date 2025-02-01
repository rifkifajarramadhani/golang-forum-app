package membership

import "github.com/gin-gonic/gin"

type Handler struct {
	*gin.Engine
}

func NewHandler(api *gin.Engine) *Handler {
	return &Handler{Engine: api}
}

func (h *Handler) Register() {
	route := h.Group("/membership")
	route.GET("/ping", h.Ping)
}