package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("record not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts
	Users
	Comments
	Followers
	Roles
}
type Posts interface {
	Create(context.Context, *Post) error
	GetByID(context.Context, int64) (*Post, error)
	Delete(context.Context, int64) error
	Update(context.Context, *Post) error
	GetUserFeed(context.Context, int64, PaginatedFeedQuery) ([]PostWithMetadata, error)
}
type Users interface {
	Create(context.Context, *sql.Tx, *User) error
	GetByID(context.Context, int64) (*User, error)
	CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error
	Activate(ctx context.Context, token string) error
	Delete(ctx context.Context, userID int64) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}
type Comments interface {
	GetByPostID(context.Context, int64) ([]Comment, error)
	Create(context.Context, *Comment) error
}
type Followers interface {
	Follow(ctx context.Context, followerID, userID int64) error
	Unfollow(ctx context.Context, followerID, userID int64) error
}

type Roles interface {
	GetByName(ctx context.Context, roleName string) (*Role, error)
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostsStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db: db},
		Followers: &FollowerStore{db: db},
		Roles:     &RoleStore{db: db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
