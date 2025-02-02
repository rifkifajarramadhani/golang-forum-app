package post

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/post"
)

func (s *service) GetPostsByUserID(ctx context.Context, userId uint64) ([]post.PostModel, error) {
	return s.postRepository.GetPostsByUserID(ctx, userId)
}
