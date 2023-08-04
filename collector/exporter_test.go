package collector

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var opts = &redis.UniversalOptions{
	Addrs:      []string{":6379"},
	Password:   "",
	ClientName: "redis_exporter_test",
}

func TestGetRedisVersion(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewUniversalClient(opts)
	v, err := getRedisVersion(ctx, rdb, nil)

	assert.Nil(t, err)
	assert.Equal(t, v, 5.0)
}
