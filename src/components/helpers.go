package cache_redis_config

import (
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/cache/v8/redis"
	redis2 "github.com/go-redis/redis/v9"
)

// RedisConfig holds configuration for a Redis instance.
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
	TLS      bool
	Cert     string
	Key      string
}

// RedisDatabaseConfig holds configuration for a Redis database instance.
type RedisDatabaseConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// NewRedisConfig returns a new RedisConfig instance from environment variables.
func NewRedisConfig() (*RedisConfig, error) {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		return nil, errors.New("REDIS_HOST is required")
	}

	portStr := os.Getenv("REDIS_PORT")
	if portStr == "" {
		return nil, errors.New("REDIS_PORT is required")
	}
	var port int
	if port, err := strconv.Atoi(portStr); err != nil {
		return nil, err
	}

	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		return nil, errors.New("REDIS_PASSWORD is required")
	}

	db := os.Getenv("REDIS_DB")
	if db == "" {
		return nil, errors.New("REDIS_DB is required")
	}
	var dbInt int
	if dbInt, err := strconv.Atoi(db); err != nil {
		return nil, err
	}

	tls := os.Getenv("REDIS_TLS") == "true"

	cert := ""
	key := ""

	return &RedisConfig{
		Host:     host,
		Port:     port,
		Password: password,
		DB:       dbInt,
		TLS:      tls,
		Cert:     cert,
		Key:      key,
	}, nil
}

// NewRedisDatabaseConfig returns a new RedisDatabaseConfig instance from environment variables.
func NewRedisDatabaseConfig() (*RedisDatabaseConfig, error) {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		return nil, errors.New("REDIS_HOST is required")
	}

	portStr := os.Getenv("REDIS_PORT")
	if portStr == "" {
		return nil, errors.New("REDIS_PORT is required")
	}
	var port int
	if port, err := strconv.Atoi(portStr); err != nil {
		return nil, err
	}

	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		return nil, errors.New("REDIS_PASSWORD is required")
	}

	db := os.Getenv("REDIS_DB")
	if db == "" {
		return nil, errors.New("REDIS_DB is required")
	}
	var dbInt int
	if dbInt, err := strconv.Atoi(db); err != nil {
		return nil, err
	}

	return &RedisDatabaseConfig{
		Host:     host,
		Port:     port,
		Password: password,
		DB:       dbInt,
	}, nil
}

// NewRedisClient returns a new Redis client instance from a RedisConfig.
func NewRedisClient(config *RedisConfig) (*cache.RedisCache, error) {
	addr := strings.Join([]string{config.Host, strconv.Itoa(config.Port)}, ":")
	if config.TLS {
		return cache.New(&cache.Config{
			Redis: redis.New(&redis.Config{
				Addr:       addr,
				Password:    config.Password,
				DB:         config.DB,
				TLSConfig: &redis2.TLSConfig{
					 Cert: config.Cert,
					Key:  config.Key,
				},
			}),
		})
	} else {
		return cache.New(&cache.Config{
			Redis: redis.New(&redis.Config{
				Addr:       addr,
				Password:    config.Password,
				DB:         config.DB,
			}),
		})
	}
}

// NewRedisDatabaseClient returns a new Redis client instance from a RedisDatabaseConfig.
func NewRedisDatabaseClient(config *RedisDatabaseConfig) (*cache.RedisCache, error) {
	addr := strings.Join([]string{config.Host, strconv.Itoa(config.Port)}, ":")
	return cache.New(&cache.Config{
		Redis: redis.New(&redis.Config{
			Addr:       addr,
			Password:    config.Password,
			DB:         config.DB,
		}),
	})
}