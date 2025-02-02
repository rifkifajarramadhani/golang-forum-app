package auth

// Request models
type (
	RegisterRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

// DB models
type (
	UserModel struct {
		ID        uint64 `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
