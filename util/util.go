package util

import (
	"context"
	"strings"

	"github.com/redis/go-redis/v9"
)

func GetAllRedisNodes(ctx context.Context, rdb *redis.Client) []string {
	result, err := rdb.ClusterNodes(ctx).Result()
	defer rdb.Close()

	if err != nil {
		return []string{}
	}

	var addrs []string
	result = strings.TrimSpace(result)
	lines := strings.Split(result, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		tokens := strings.Split(line, " ")
		addrs = append(addrs, strings.Split(tokens[1], "@")[0])
	}

	return addrs
}

func ParseInfoSection(section string) map[string]string {
	section = strings.TrimSpace(section)
	lines := strings.Split(section, "\n")

	m := make(map[string]string, 4)
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			arr := strings.Split(line, ":")
			m[strings.TrimSpace(arr[0])] = strings.TrimSpace(arr[1])
		}
	}

	return m
}
