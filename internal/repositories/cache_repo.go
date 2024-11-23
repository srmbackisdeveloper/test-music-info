package repositories

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheRepository struct {
	Client *redis.Client
}

func NewRedisClient(addr, password string, db int) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })

    if _, err := rdb.Ping(context.Background()).Result(); err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }

    log.Println("Connected to Redis")
    return rdb
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
    return &CacheRepository{Client: client}
}

// methods:
func (repo *CacheRepository) SetSongCache(key string, value string, ttl time.Duration) error {
    ctx := context.Background()
    return repo.Client.Set(ctx, key, value, ttl).Err()
}

func (repo *CacheRepository) GetSongCache(key string) (string, error) {
    ctx := context.Background()
    result, err := repo.Client.Get(ctx, key).Result()
	if err == redis.Nil {
        return "", nil
    } else if err != nil {
        return "", err
    }
    return result, nil
}

func (repo *CacheRepository) DeleteSongCache(key string) error {
    ctx := context.Background()
    return repo.Client.Del(ctx, key).Err()
}