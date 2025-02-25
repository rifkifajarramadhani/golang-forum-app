package post

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/post"
)

func (s *Service) GetPosts(ctx context.Context) ([]post.PostModel, error) {
	return s.postRepository.GetPosts(ctx)
}
