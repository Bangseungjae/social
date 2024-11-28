package cache

import (
	"context"
	"github.com/Bangseungjae/social/internal/store"
	"github.com/go-redis/redis/v8"
)

type Storage struct {
	Users
}

type Users interface {
	Get(ctx context.Context, userID int64) (*store.User, error)
	Set(ctx context.Context, user *store.User) error
	Delete(ctx context.Context, userID int64)
}

func NewRedisStore(redisDB *redis.Client) Storage {
	return Storage{
		Users: &UsersStore{redisDB: redisDB},
	}
}
