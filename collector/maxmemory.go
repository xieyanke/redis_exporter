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
	"strconv"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	redis "github.com/redis/go-redis/v9"
)

const (
	maxMemory = "maxmemory"
)

var (
	maxMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, maxMemory, "bytes"),
		"Collect redis config maxmemory.",
		nil,
		nil,
	)
)

type MaxMemoryScraper struct{}

// Help implements Scraper.
func (MaxMemoryScraper) Help() string {
	return "Collect max memory from redis"
}

// Name implements Scraper.
func (MaxMemoryScraper) Name() string {
	return "maxmemroy"
}

// Scrape implements Scraper.
func (MaxMemoryScraper) Scrape(ctx context.Context, rdb redis.UniversalClient, ch chan<- prometheus.Metric, logger log.Logger) error {
	resultMap, err := rdb.ConfigGet(ctx, "maxmemory").Result()
	if err != nil {
		return err
	}

	maxMemoryInBytes, err := strconv.ParseFloat(resultMap["maxmemory"], 64)
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(
		maxMemoryDesc,
		prometheus.GaugeValue,
		maxMemoryInBytes,
	)

	return err
}

// Version implements Scraper.
func (MaxMemoryScraper) Version() string {
	return "2.4"
}

var _ Scraper = MaxMemoryScraper{}
