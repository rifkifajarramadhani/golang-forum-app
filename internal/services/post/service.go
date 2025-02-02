package post

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/post"
)

type postRepositoryInterface interface {
	CreatePost(ctx context.Context, model post.PostModel) error
	GetPost(ctx context.Context, id uint64) (*post.PostModel, error)
	GetPosts(ctx context.Context) ([]post.PostModel, error)
	GetPostsByUserID(ctx context.Context, userId uint64) ([]post.PostModel, error)
	UpdatePost(ctx context.Context, id uint64, model post.PostModel) error
}

type service struct {
	postRepository postRepositoryInterface
}

func NewService(postRepository postRepositoryInterface) *service {
	return &service{
		postRepository: postRepository,
	}
}
