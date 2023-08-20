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
	"sync"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/redis/go-redis/v9"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace = "redis"
	subsystem = "exporter"
)

var (
	redisUp = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "", "up"),
		"Whether the redis service is up.",
		nil,
		nil,
	)

	redisScrapeDurationSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, subsystem, "scrape_duration_seconds"),
		"Collector scrape duration.",
		[]string{"collector"},
		nil,
	)

	redisScrapeSuccess = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, subsystem, "scrape_success"),
		"redis_exporter: Whether a collector scrape success.",
		[]string{"collector"},
		nil,
	)
)

type MetricDesc struct {
	Subsystem string
	Name      string
	Help      string
	Labels    []string
}

// Exporter collects redis metrics.
type Exporter struct {
	ctx      context.Context
	logger   log.Logger
	opts     []*redis.Options
	scrapers []Scraper
}

// Collect implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.scrape(e.ctx, ch)
	ch <- prometheus.MustNewConstMetric(redisUp, prometheus.CounterValue, 1)
}

// Describe implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- redisUp
	ch <- redisScrapeDurationSeconds
}

// *Exporter implements prometheus.Collector
var _ prometheus.Collector = (*Exporter)(nil)

func (e *Exporter) scrape(ctx context.Context, ch chan<- prometheus.Metric) float64 {
	var rdbs []*redis.Client
	for _, opt := range e.opts {
		rdbs = append(rdbs, redis.NewClient(opt))
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	for _, scraper := range e.scrapers {
		wg.Add(1)
		go func(scraper Scraper) {
			defer wg.Done()

			scrapeSuccess := 1.0
			label := fmt.Sprintf("collect.%s", scraper.Name())
			startTime := time.Now()
			if err := scraper.Scrape(ctx, rdbs, ch, log.With(e.logger, "scraper", scraper.Name())); err != nil {
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

func New(ctx context.Context, opts []*redis.Options, scrapers []Scraper, logger log.Logger) *Exporter {
	return &Exporter{
		ctx:      ctx,
		logger:   logger,
		opts:     opts,
		scrapers: scrapers,
	}
}
