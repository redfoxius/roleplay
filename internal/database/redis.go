package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisDB represents a Redis database connection
type RedisDB struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisDB creates a new Redis database connection
func NewRedisDB(addr, password string, db int) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()

	// Test connection
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &RedisDB{
		client: client,
		ctx:    ctx,
	}, nil
}

// Close closes the Redis connection
func (db *RedisDB) Close() error {
	return db.client.Close()
}

// Set stores a value in Redis with an expiration time
func (db *RedisDB) Set(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %v", err)
	}

	return db.client.Set(db.ctx, key, data, expiration).Err()
}

// Get retrieves a value from Redis
func (db *RedisDB) Get(key string, value interface{}) error {
	data, err := db.client.Get(db.ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("key not found: %s", key)
		}
		return fmt.Errorf("failed to get value: %v", err)
	}

	return json.Unmarshal(data, value)
}

// Delete removes a value from Redis
func (db *RedisDB) Delete(key string) error {
	return db.client.Del(db.ctx, key).Err()
}

// Exists checks if a key exists in Redis
func (db *RedisDB) Exists(key string) (bool, error) {
	n, err := db.client.Exists(db.ctx, key).Result()
	return n > 0, err
}

// Keys returns all keys matching a pattern
func (db *RedisDB) Keys(pattern string) ([]string, error) {
	return db.client.Keys(db.ctx, pattern).Result()
}

// SetNX sets a value only if the key doesn't exist
func (db *RedisDB) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return false, fmt.Errorf("failed to marshal value: %v", err)
	}

	return db.client.SetNX(db.ctx, key, data, expiration).Result()
}

// HSet sets a field in a hash
func (db *RedisDB) HSet(key string, field string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %v", err)
	}

	return db.client.HSet(db.ctx, key, field, data).Err()
}

// HGet gets a field from a hash
func (db *RedisDB) HGet(key string, field string, value interface{}) error {
	data, err := db.client.HGet(db.ctx, key, field).Bytes()
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("field not found: %s", field)
		}
		return fmt.Errorf("failed to get field: %v", err)
	}

	return json.Unmarshal(data, value)
}

// HGetAll gets all fields from a hash
func (db *RedisDB) HGetAll(key string) (map[string]string, error) {
	return db.client.HGetAll(db.ctx, key).Result()
}

// HDel deletes a field from a hash
func (db *RedisDB) HDel(key string, field string) error {
	return db.client.HDel(db.ctx, key, field).Err()
}

// SAdd adds members to a set
func (db *RedisDB) SAdd(key string, members ...interface{}) error {
	return db.client.SAdd(db.ctx, key, members...).Err()
}

// SMembers gets all members of a set
func (db *RedisDB) SMembers(key string) ([]string, error) {
	return db.client.SMembers(db.ctx, key).Result()
}

// SRem removes members from a set
func (db *RedisDB) SRem(key string, members ...interface{}) error {
	return db.client.SRem(db.ctx, key, members...).Err()
}
