package helper

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetOrSetCache[T any](redisClient *redis.Client, cacheKey string, ttl time.Duration, fetchData func() (T, error)) (T, error) {
	ctx := context.Background()

	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var result T
		json.Unmarshal([]byte(cachedData), &result)
		log.Println("Cache hit:", cacheKey)
		return result, nil
	}

	log.Println("Cache miss:", cacheKey)
	data, err := fetchData()
	if err != nil {
		return data, err
	}

	dataJSON, _ := json.Marshal(data)
	err = redisClient.Set(ctx, cacheKey, dataJSON, ttl).Err()
	if err != nil {
		log.Println("Gagal menyimpan data ke Redis:", err)
	}

	return data, nil
}
