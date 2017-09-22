package redis

import (
	r "github.com/go-redis/redis"
)

type ClientRedis struct {
	redisClient *r.Client
}

func NewClient(address string, password string, dbNum int) *ClientRedis {
	client := new(ClientRedis)
	client.redisClient = r.NewClient(&r.Options{
		Addr:     address,
		Password: password,
		DB:       dbNum,
	})

	return client
}

func (c *ClientRedis) Keys(basKey string) ([]string, error) {
	keys, err := c.redisClient.Keys(basKey + "|*").Result()
	if err != nil {
		return []string{}, err
	}
	return keys, nil
}

func (c *ClientRedis) Set(key, val string) error {
	return c.redisClient.Set(key, val, 0).Err()
}

func (c *ClientRedis) Del(keys ...string) error {
	return c.redisClient.Del(keys...).Err()
}
