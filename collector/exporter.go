/*
Copyright 2023 XieYanke.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package collector

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/redis/go-redis/v9"
	"github.com/xieyanke/redis_exporter/util"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "redis"
	subsystem = "exporter"
)

var (
	redisUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Whether the redis node is up.",
		nil,
		nil,
	)

	redisNodeUptime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "uptime_seconds"),
		"Number of seconds since the redis node started.",
		[]string{"node_id"},
		nil,
	)

	redisScrapeDurationSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, "scrape_duration_seconds"),
		"Collector scrape duration.",
		[]string{"collector"},
		nil,
	)

	redisScrapeSuccess = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, "scrape_success"),
		"redis_exporter: Whether a collector scrape success.",
		[]string{"collector"},
		nil,
	)
)

// Exporter collects redis metrics.
type Exporter struct {
	ctx      context.Context
	logger   log.Logger
	opts     *redis.UniversalOptions
	scrapers []Scraper
}

// Collect implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	up := e.scrape(e.ctx, ch)
	ch <- prometheus.MustNewConstMetric(redisUp, prometheus.CounterValue, up)
}

// Describe implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- redisUp
	ch <- redisScrapeDurationSeconds
	ch <- redisNodeUptime
}

// *Exporter implements prometheus.Collector
var _ prometheus.Collector = (*Exporter)(nil)

func (e *Exporter) scrape(ctx context.Context, ch chan<- prometheus.Metric) float64 {
	var err error
	scrapeTime := time.Now()

	rdb := redis.NewUniversalClient(e.opts)

	defer rdb.Close()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		level.Error(e.logger).Log("msg", "Error pinging redis", "err", err)
		return 0.0
	}

	ch <- prometheus.MustNewConstMetric(redisScrapeDurationSeconds, prometheus.GaugeValue, time.Since(scrapeTime).Seconds(), "connection")

	var wg sync.WaitGroup
	defer wg.Wait()

	for _, scraper := range e.scrapers {
		wg.Add(1)
		go func(scraper Scraper) {
			defer wg.Done()

			scrapeSuccess := 1.0
			label := fmt.Sprintf("collect.%s", scraper.Name())
			startTime := time.Now()
			if err := scraper.Scrape(ctx, rdb, ch, log.With(e.logger, "scraper", scraper.Name())); err != nil {
				level.Error(e.logger).Log("msg", "Error from scraper", "scraper", scraper.Name(), "err", err)
				scrapeSuccess = 0.0
			}
			duration := time.Since(startTime).Seconds()
			ch <- prometheus.MustNewConstMetric(redisScrapeDurationSeconds, prometheus.GaugeValue, duration, label)
			ch <- prometheus.MustNewConstMetric(redisScrapeSuccess, prometheus.GaugeValue, scrapeSuccess, label)
		}(scraper)
	}

	return 1.0
}

func New(ctx context.Context, opts *redis.UniversalOptions, scrapers []Scraper, logger log.Logger) *Exporter {
	return &Exporter{
		ctx:      ctx,
		logger:   logger,
		opts:     opts,
		scrapers: scrapers,
	}
}

var versionRE = regexp.MustCompile(`^\d+\.\d+`)

func getRedisVersion(ctx context.Context, rdb redis.UniversalClient, logger log.Logger) (float64, error) {
	var versionStr string
	var versionNum float64

	section, err := rdb.Info(ctx, "Server").Result()
	if err != nil {
		return -1.0, err
	}

	version := util.ParseInfoSection(section)["redis_version"]
	versionStr = versionRE.FindString(version)
	versionNum, err = strconv.ParseFloat(versionStr, 64)
	if err != nil {
		return -1.0, err
	}

	return versionNum, nil
}
