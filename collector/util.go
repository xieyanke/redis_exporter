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

func parseRedisInfoResp(resp string) map[string]string {
	resp = strings.TrimSpace(resp)
	lines := strings.Split(resp, "\n")

	m := make(map[string]string, 8)
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			lineArr := strings.Split(line, ":")
			key := strings.TrimSpace(lineArr[0])
			value := strings.TrimSpace(lineArr[1])
			switch value {
			case "ok", "up":
				value = "1"
			case "down":
				value = "0"
			}

			m[key] = value
		}
	}
	return m
}

func parseRedisInfoKeyspaceOrCmdtatsResp(resp string) map[string]string {
	resp = strings.TrimSpace(resp)
	lines := strings.Split(resp, "\n")

	m := make(map[string]string, 8)
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			lineArr := strings.Split(line, ":")
			prefix := strings.TrimSpace(lineArr[0])
			items := strings.Split(strings.TrimSpace(lineArr[1]), ",")
			for _, item := range items {
				itemArr := strings.Split(item, "=")
				m[fmt.Sprintf("%s_%s", prefix, itemArr[0])] = itemArr[1]
			}
		}
	}

	return m
}

func checkParseRedisInfoRespError(key, addr string, err error, logger log.Logger) {
	if err != nil {
		level.Error(logger).Log("msg", fmt.Sprintf("parse %s value failed from %s", key, addr), "err", err)
	}
}
