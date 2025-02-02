package post

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/post"
)

func (s *service) CreatePost(ctx context.Context, req post.CreatePostRequest) error {
	model := post.PostModel{
		UserID:  req.UserID,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := s.postRepository.CreatePost(ctx, model); err != nil {
		return err
	}

	return nil
}
