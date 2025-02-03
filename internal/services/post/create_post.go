package post

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/post"
)

func (s *Service) CreatePost(ctx context.Context, userId uint64, req post.CreatePostRequest) error {
	model := post.PostModel{
		UserID:  userId,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := s.postRepository.CreatePost(ctx, model); err != nil {
		return err
	}

	return nil
}
