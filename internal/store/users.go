package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	IsActive  bool   `json:"is_active"`
}

type UsersRepository struct {
	db *sql.DB
}

func (r *UsersRepository) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, created_at`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
