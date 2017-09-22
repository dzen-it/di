package cache

import (
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

type ClientCache struct {
	cacheClient *cache.Cache
}

func NewClient(est int) *ClientCache {
	client := new(ClientCache)
	client.cacheClient = cache.New(time.Second*time.Duration(est), 1)
	return client
}

func (c *ClientCache) Keys(baseKey string) ([]string, error) {
	items := c.cacheClient.Items()
	keys := make([]string, 0, len(items))
	for k := range items {
		if strings.Contains(k, baseKey+"|") {
			keys = append(keys, k)
		}
	}
	return keys, nil
}

func (c *ClientCache) Set(key, val string) error {
	c.cacheClient.Set(key, val, 0)
	return nil
}

func (c *ClientCache) Del(keys ...string) error {
	for i := range keys {
		c.cacheClient.Delete(keys[i])
	}
	return nil
}
