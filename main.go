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

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
	"github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"github.com/redis/go-redis/v9"
	"github.com/xieyanke/redis_exporter/collector"
)

var (
	webConfig          = kingpinflag.AddFlags(kingpin.CommandLine, ":9121")
	metricsPath        = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
	addrs              = kingpin.Flag("redis.addrs", "Redis server addresses.").Default("localhost:6379").Strings()
	passwd             = kingpin.Flag("redis.passwd", "Redis server password.").Default("").String()
	db                 = kingpin.Flag("redis.db", "Redis db number.").Default("0").Int()
	mode               = kingpin.Flag("redis.mode", "Redis server mode.").Default("standalone").String()
	clientName         = kingpin.Flag("redis.client-name", "Redis client name.").Default("redis_exporter").String()
	keyFile            = kingpin.Flag("redis.tls.key-file", "Client private key file.").Default("").String()
	caFile             = kingpin.Flag("redis.tls.ca-file", "Client root ca file.").Default("").String()
	insecureSkipVerify = kingpin.Flag("redis.tls.insecure-skip-verify", "Skip server certificate verification.").Bool()
	timeout            = kingpin.Flag("redis.timeout", "Redis connect timeout.").Default("1s").Duration()
)

func init() {
	prometheus.MustRegister(version.NewCollector("redis_exporter"))
}

var scrapersTable = map[collector.Scraper]bool{
	collector.NewInfoClientsScraper():      true,
	collector.NewInfoCPUScraper():          true,
	collector.NewInfoServerScraper():       true,
	collector.NewInfoMemoryScraper():       true,
	collector.NewInfoReplicationScraper():  true,
	collector.NewInfoPersistenceScraper():  true,
	collector.NewInfoStatsScraper():        true,
	collector.NewInfoKeyspaceScraper():     true,
	collector.NewInfoCommandStatsScraper(): true,
}

func newHandler(scrapers []collector.Scraper, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var timeoutSeconds float64
		if v := r.Header.Get("X-Prometheus-Scrape-Timeout-Seconds"); v != "" {
			timeoutSeconds, err = strconv.ParseFloat(v, 64)
			if err != nil {
				level.Error(logger).Log("msg", "Failed to parse timeout from Prometheus Header", "err", err)
			}
		}

		if timeoutSeconds > 0 {
			*timeout = time.Duration(timeoutSeconds * float64(time.Second))
		}

		ctx := r.Context()

		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()

		r = r.WithContext(ctx)
		var initCli *redis.Client
		for _, addr := range *addrs {
			initCli = redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: *passwd,
			})

			err = initCli.Ping(ctx).Err()
			if err == nil {
				break
			} else {
				level.Error(logger).Log("msg", fmt.Sprintf("%s can't connect", addr), "err", err)
			}
		}
		defer initCli.Close()

		var allAddrs []string
		switch *mode {
		case "cluster":
			allAddrs, _ = collector.GetRedisClusterNodes(ctx, initCli)
		default:
			allAddrs = *addrs
		}

		var opts []*redis.Options
		for _, addr := range allAddrs {
			opts = append(opts, &redis.Options{Addr: addr, Password: *passwd})
		}

		registry := prometheus.NewRegistry()
		registry.MustRegister(collector.New(ctx, opts, scrapers, logger))
		gatherers := prometheus.Gatherers{
			prometheus.DefaultGatherer,
			registry,
		}

		h := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)
	}
}

func main() {
	// Generate ON/OFF flags for all scrapers.
	scraperFlags := map[collector.Scraper]*bool{}
	for scraper, enabledByDefault := range scrapersTable {
		defaultOn := "false"
		if enabledByDefault {
			defaultOn = "true"
		}

		f := kingpin.Flag(
			fmt.Sprintf("collect.%s", scraper.Name()),
			scraper.Help(),
		).Default(defaultOn).Bool()

		scraperFlags[scraper] = f
	}

	promlogconfig := &promlog.Config{}

	flag.AddFlags(kingpin.CommandLine, promlogconfig)
	kingpin.HelpFlag.Short('h')
	kingpin.Version(version.Print("redis_exporter"))
	kingpin.Parse()

	logger := promlog.New(promlogconfig)

	enabledScrapers := []collector.Scraper{}

	for scraper, enabled := range scraperFlags {
		if *enabled {
			level.Info(logger).Log("msg", "Scraper enabled", "scraper", scraper.Name())
			enabledScrapers = append(enabledScrapers, scraper)
		}
	}

	if *mode == "cluster" {
		scraper := collector.NewClusterInfoScraper()
		enabledScrapers = append(enabledScrapers, scraper)
	}

	handlerFunc := newHandler(enabledScrapers, logger)
	http.Handle(*metricsPath, promhttp.InstrumentMetricHandler(prometheus.DefaultRegisterer, handlerFunc))

	if *metricsPath != "/" && *metricsPath != "" {
		landingConfig := web.LandingConfig{
			Name:        "Redis Exporter",
			Description: "Prometheus Exporter for redis service",
			Version:     version.Info(),
			Links: []web.LandingLinks{
				{
					Address: *metricsPath,
					Text:    "Metrics",
				},
			},
		}

		landingPage, err := web.NewLandingPage(landingConfig)
		if err != nil {
			level.Error(logger).Log("err", err)
			os.Exit(1)
		}

		http.Handle("/", landingPage)
	}

	srv := &http.Server{}
	level.Info(logger).Log("msg", "Starting redis_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "context", version.BuildContext())

	if err := web.ListenAndServe(srv, webConfig, logger); err != nil {
		level.Error(logger).Log("msg", "Error starting HTTP Server", "err", err)
		os.Exit(1)
	}
}
