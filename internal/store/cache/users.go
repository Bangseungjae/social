package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Bangseungjae/social/internal/store"
	"github.com/go-redis/redis/v8"
	"time"
)

type UsersStore struct {
	redisDB *redis.Client
}

const UserExpTime = time.Minute

func (u *UsersStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	cacheKey := fmt.Sprintf("user-%d", userID)

	data, err := u.redisDB.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user store.User
	if data != "" {
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UsersStore) Set(ctx context.Context, user *store.User) error {
	cacheKey := fmt.Sprintf("user-%v", user.ID)

	json, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// EX -> 만료시간이 있는
	return u.redisDB.SetEX(ctx, cacheKey, json, UserExpTime).Err()
}

func (s *UsersStore) Delete(ctx context.Context, userID int64) {
	cacheKey := fmt.Sprintf("user-%d", userID)
	s.redisDB.Del(ctx, cacheKey)
}
