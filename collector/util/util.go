package util

import (
	"context"
	"strings"

	"github.com/redis/go-redis/v9"
)

func GetRedisClusterNodes(ctx context.Context, rdb *redis.Client) ([]string, error) {
	result, err := rdb.ClusterNodes(ctx).Result()
	defer rdb.Close()

	if err != nil {
		return []string{}, err
	}

	var addrs []string
	result = strings.TrimSpace(result)
	lines := strings.Split(result, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		tokens := strings.Split(line, " ")
		addrs = append(addrs, strings.Split(tokens[1], "@")[0])
	}

	return addrs, nil
}

