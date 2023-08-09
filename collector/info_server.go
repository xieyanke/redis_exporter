package collector

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	redis "github.com/redis/go-redis/v9"
	"github.com/xieyanke/redis_exporter/util"
)

var (
	redisNodeServerUptimeSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "server", "node_uptime_seconds"),
		"Collect redis uptime in seconds and redis server node basic infos.",
		[]string{"addr", "version", "redis_build_id", "mode", "os", "multiplexing_api", "atomicvar_api", "gcc_version"},
		nil,
	)
)

type InfoServerScraper struct {
}

// Help implements Scraper.
func (*InfoServerScraper) Help() string {
	return "Collect redis node info."
}

// Name implements Scraper.
func (*InfoServerScraper) Name() string {
	return "info server"
}

// Scrape implements Scraper.
func (*InfoServerScraper) Scrape(ctx context.Context, rdbs []*redis.Client, ch chan<- prometheus.Metric, logger log.Logger) error {
	var err error
	for _, rdb := range rdbs {
		addr := rdb.Options().Addr
		section, err := rdb.Info(ctx, "server").Result()
		if err != nil {
			return err
		}

		sectionMap := util.ParseInfoSection(section)
		uptimeFloat64, err := strconv.ParseFloat(sectionMap["uptime_in_seconds"], 64)
		checkParseResultError("uptime_in_seconds", addr, err, logger)
		ch <- prometheus.MustNewConstMetric(
			redisNodeServerUptimeSeconds,
			prometheus.GaugeValue,
			uptimeFloat64,
			addr, sectionMap["redis_version"], sectionMap["redis_build_id"],
			sectionMap["redis_mode"], sectionMap["os"], sectionMap["multiplexing_api"],
			sectionMap["atomicvar_api"], sectionMap["gcc_version"],
		)
	}
	return err
}

// Version implements Scraper.
func (*InfoServerScraper) Version() string {
	return "1.0"
}

var _ Scraper = &InfoServerScraper{}

func GetRedisMode(ctx context.Context, rdb *redis.Client, logger log.Logger) string {
	section, err := rdb.Info(ctx, "server").Result()
	if err != nil {
		level.Error(logger).Log("msg", fmt.Sprintf("fail to get redis mode from %s", rdb.Options().Addr), "err", err)
	}

	sectionMap := util.ParseInfoSection(section)

	return sectionMap["redis_mode"]
}
