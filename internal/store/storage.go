package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
}

func NewRepository(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsRepository{db},
		Users: &UsersRepository{db},
	}
}
