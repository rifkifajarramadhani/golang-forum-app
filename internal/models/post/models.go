package post

// Request models
type (
	CreatePostRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	UpdatePostRequest struct {
		ID      uint64 `json:"id"`
		UserID  uint64 `json:"user_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
)

// Response models
type (
	GetPostsResponse struct {
		Posts      []PostModel `json:"posts"`
		Pagination Pagination  `json:"pagination"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
)

// DB models
type (
	PostModel struct {
		ID        uint64 `json:"id"`
		UserID    uint64 `json:"user_id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
