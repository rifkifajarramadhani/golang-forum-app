package post

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/post"
)

func (s *service) GetPost(ctx context.Context, id uint64) (*post.PostModel, error) {
	return s.postRepository.GetPost(ctx, id)
}
