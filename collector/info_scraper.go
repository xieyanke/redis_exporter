package collector

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	redis "github.com/redis/go-redis/v9"
)

type InfoScraper struct {
	section     string
	sectionHelp string
	metricsDesc map[string]*MetricDesc
}

// Help implements Scraper.
func (scraper *InfoScraper) Help() string {
	return scraper.sectionHelp
}

// Name implements Scraper.
func (scraper *InfoScraper) Name() string {
	return fmt.Sprintf("info.%s", scraper.section)
}

// Version implements Scraper.
func (*InfoScraper) Version() string {
	return "1.0"
}

// Scrape implements Scraper.
func (scraper *InfoScraper) Scrape(ctx context.Context, rdbs []*redis.Client, ch chan<- prometheus.Metric, logger log.Logger) error {
	var err error

	for _, rdb := range rdbs {
		addr := rdb.Options().Addr

		var sectionRes string
		sectionRes, err = rdb.Info(ctx, scraper.section).Result()
		if err != nil {
			return err
		}

		sectionMap := parseInfoSection(sectionRes)

		for k, v := range scraper.metricsDesc {
			var f64 float64
			f64, err = strconv.ParseFloat(sectionMap[k], 64)
			checkParseResultError(k, addr, err, logger)

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

var _ Scraper = &InfoScraper{}

func parseInfoSection(section string) map[string]string {
	section = strings.TrimSpace(section)
	lines := strings.Split(section, "\n")

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
