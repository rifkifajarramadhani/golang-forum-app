package post

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/post"
)

type PostRepositoryInterface interface {
	CreatePost(ctx context.Context, model post.PostModel) error
	GetPost(ctx context.Context, id uint64) (*post.PostModel, error)
	GetPosts(ctx context.Context) ([]post.PostModel, error)
	GetPostsByUserID(ctx context.Context, userId uint64) ([]post.PostModel, error)
	UpdatePost(ctx context.Context, id uint64, model post.PostModel) error
}

type Service struct {
	postRepository PostRepositoryInterface
}

func NewService(postRepository PostRepositoryInterface) *Service {
	return &Service{
		postRepository: postRepository,
	}
}
