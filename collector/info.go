package collector

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	redis "github.com/redis/go-redis/v9"
)

type infoScraper struct {
	section     string
	sectionHelp string
	metricsDesc map[string]*MetricDesc
}

// Help implements Scraper.
func (scraper *infoScraper) Help() string {
	return scraper.sectionHelp
}

// Name implements Scraper.
func (scraper *infoScraper) Name() string {
	return fmt.Sprintf("info.%s", scraper.section)
}

// Version implements Scraper.
func (*infoScraper) Version() string {
	return "1.0"
}

// Scrape implements Scraper.
func (scraper *infoScraper) Scrape(ctx context.Context, rdbs []*redis.Client, ch chan<- prometheus.Metric, logger log.Logger) error {
	var err error

	for _, rdb := range rdbs {
		addr := rdb.Options().Addr

		var sectionRes string
		sectionRes, err = rdb.Info(ctx, scraper.section).Result()
		if err != nil {
			return err
		}
		var sectionMap map[string]string

		switch scraper.section {
		case "keyspace":
			sectionMap = parseRedisInfoKeyspaceOrCmdtatsResp(sectionRes)
			metricsDesc := initKeyspaceMetricsDesc(sectionMap)
			scraper.metricsDesc = metricsDesc
		case "commandstats":
			sectionMap = parseRedisInfoKeyspaceOrCmdtatsResp(sectionRes)
			metricsDesc := initCmdStatsMetricsDesc(sectionMap)
			scraper.metricsDesc = metricsDesc
		default:
			sectionMap = parseRedisInfoResp(sectionRes)
		}

		for k, v := range scraper.metricsDesc {
			var f64 float64
			f64, err = strconv.ParseFloat(sectionMap[k], 64)
			checkParseRedisInfoRespError(k, addr, err, logger)

			desc := prometheus.NewDesc(
				prometheus.BuildFQName(Namespace, v.Subsystem, v.Name),
				v.Help,
				v.Labels,
				nil,
			)
			ch <- prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				f64,
				addr,
			)
		}
	}

	return err
}

var _ Scraper = &infoScraper{}
