package post

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/post"
)

func (s *Service) UpdatePost(ctx context.Context, id uint64, req post.UpdatePostRequest) error {
	model := post.PostModel{
		Title:   req.Title,
		Content: req.Content,
	}

	return s.postRepository.UpdatePost(ctx, id, model)
}
