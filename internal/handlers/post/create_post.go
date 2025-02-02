package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/models/post"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var req post.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.postService.CreatePost(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post created"})
}
