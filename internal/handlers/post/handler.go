package post

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/internal/middleware"
	"github.com/rifkifajarramadhani/internal/models/post"
)

type postServiceInterface interface {
	CreatePost(ctx context.Context, req post.CreatePostRequest) error
	GetPost(ctx context.Context, id uint64) (*post.PostModel, error)
	GetPosts(ctx context.Context) ([]post.PostModel, error)
	GetPostsByUserID(ctx context.Context, userId uint64) ([]post.PostModel, error)
	UpdatePost(ctx context.Context, id uint64, req post.UpdatePostRequest) error
}

type Handler struct {
	*gin.Engine

	postService postServiceInterface
}

func NewHandler(api *gin.Engine, postService postServiceInterface) *Handler {
	return &Handler{
		Engine:      api,
		postService: postService,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/posts")
	route.Use(middleware.AuthMiddleware())

	route.GET("/", h.GetPosts)
}
