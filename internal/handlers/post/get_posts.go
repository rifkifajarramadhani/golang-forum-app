package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPosts(c *gin.Context) {
	posts, err := h.postService.GetPosts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}
