package dao

import (
	"context"
	"fmt"
	"time"
)

func (d *Dao) ensKeyRedis(key string) string {
	return fmt.Sprintf("%sens_%s", d.c.Redis.Prefix, key)
}

// SetEns 保存Ens信息
func (d *Dao) SetEns(c context.Context, key, value string, expiration time.Duration) (err error) {
	return d.redis.Set(c, d.ensKeyRedis(key), value, expiration).Err()
}

// GetEns 获取Ens信息
func (d *Dao) GetEns(c context.Context, key string) (value string, err error) {
	value, err = d.redis.Get(c, d.ensKeyRedis(key)).Result()
	return
}
