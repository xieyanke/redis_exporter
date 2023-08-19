package collector

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
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

func parseRedisRespInfo(resp string) map[string]string {
	resp = strings.TrimSpace(resp)
	lines := strings.Split(resp, "\n")

	m := make(map[string]string, 4)
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			arr := strings.Split(line, ":")
			v := strings.TrimSpace(arr[1])
			switch v {
			case "ok", "up":
				v = "1"
			case "down":
				v = "0"
			}

			m[strings.TrimSpace(arr[0])] = v
		}
	}
	return m
}

func checkParseRedisRespInfoError(key, addr string, err error, logger log.Logger) {
	if err != nil {
		level.Error(logger).Log("msg", fmt.Sprintf("parse %s value failed from %s", key, addr), "err", err)
	}
}
