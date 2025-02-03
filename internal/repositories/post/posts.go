package post

import (
	"context"

	"github.com/rifkifajarramadhani/internal/models/post"
)

func (r *Repository) CreatePost(ctx context.Context, model post.PostModel) error {
	query := `INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.Title, model.Content, model.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetPost(ctx context.Context, id uint64) (*post.PostModel, error) {
	query := `SELECT id, user_id, title, content, created_at, updated_at FROM posts WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var model post.PostModel
	err := row.Scan(&model.ID, &model.UserID, &model.Title, &model.Content, &model.CreatedAt, &model.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *Repository) GetPosts(ctx context.Context) ([]post.PostModel, error) {
	query := `SELECT id, user_id, title, content, created_at, updated_at FROM posts p JOIN users u ON p.user_id = u.id LIMIT ? OFFSET ?`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []post.PostModel
	for rows.Next() {
		var model post.PostModel
		err := rows.Scan(&model.ID, &model.UserID, &model.Title, &model.Content, &model.CreatedAt, &model.UpdatedAt)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}
	return models, nil
}

func (r *Repository) GetPostsByUserID(ctx context.Context, userID uint64) ([]post.PostModel, error) {
	query := `SELECT id, user_id, title, content, created_at, updated_at FROM posts WHERE user_id = ?`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []post.PostModel
	for rows.Next() {
		var model post.PostModel
		err := rows.Scan(&model.ID, &model.UserID, &model.Title, &model.Content, &model.CreatedAt, &model.UpdatedAt)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}
	return models, nil
}

func (r *Repository) UpdatePost(ctx context.Context, id uint64, model post.PostModel) error {
	query := `UPDATE posts SET title = ?, content = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, model.Title, model.Content, id)
	if err != nil {
		return err
	}
	return nil
}
