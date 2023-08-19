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

// sharding cluster metrics
var clusterInfoMetricsDesc = map[string]*MetricDesc{
	"cluster_state": &MetricDesc{
		Subsystem: "cluster",
		Name:      "state",
		Help:      "Flag redis cluster is active.",
		Labels:    []string{"addr"},
	},
	"cluster_slots_assigned": &MetricDesc{
		Subsystem: "cluster",
		Name:      "slots_assigned_in_total",
		Help:      "Number of redis cluster slots assinged.",
		Labels:    []string{"addr"},
	},
	"cluster_slots_ok": &MetricDesc{
		Subsystem: "cluster",
		Name:      "slots_ok_in_total",
		Help:      "Number of redis cluster healthy slots.",
		Labels:    []string{"addr"},
	},
	"cluster_slots_pfail": &MetricDesc{
		Subsystem: "cluster",
		Name:      "slots_pfail_in_total",
		Help:      "Number of redis cluster probably failed slots.",
		Labels:    []string{"addr"},
	},
	"cluster_slots_fail": &MetricDesc{
		Subsystem: "cluster",
		Name:      "slots_fail_in_total",
		Help:      "Number of redis cluster failed slots.",
		Labels:    []string{"addr"},
	},
	"cluster_known_nodes": &MetricDesc{
		Subsystem: "cluster",
		Name:      "known_nodes_in_total",
		Help:      "Number of all the redis cluster nodes.",
		Labels:    []string{"addr"},
	},
	"cluster_size": &MetricDesc{
		Subsystem: "cluster",
		Name:      "master_nodes_in_total",
		Help:      "Number of all the master nodes.",
		Labels:    []string{"addr"},
	},
	"cluster_current_epoch": &MetricDesc{
		Subsystem: "cluster",
		Name:      "current_epoch_count",
		Help:      "The current epoch of redis cluster",
		Labels:    []string{"addr"},
	},
	"cluster_my_epoch": &MetricDesc{
		Subsystem: "cluster",
		Name:      "my_epoch_count",
		Help:      "The current epoch of the current node.",
		Labels:    []string{"addr"},
	},
	"cluster_stats_messages_ping_sent": &MetricDesc{
		Subsystem: "cluseter",
		Name:      "stats_messages_ping_sent_in_bytes",
		Help:      "Total number of bytes that cluster stats messages ping sent.",
		Labels:    []string{"addr"},
	},
	"cluster_stats_messages_pong_sent": &MetricDesc{
		Subsystem: "cluster",
		Name:      "stats_messages_pong_sent_in_bytes",
		Help:      "Total number of bytes that cluster stats messages pong sent.",
		Labels:    []string{"addr"},
	},
	"cluster_stats_messages_publish_sent": &MetricDesc{
		Subsystem: "cluster",
		Name:      "stats_messages_publish_sent_in_bytes",
		Help:      "Total number of bytes that cluster stats messages publish sent.",
		Labels:    []string{"addr"},
	},
	"cluster_stats_messages_sent": &MetricDesc{
		Subsystem: "cluster",
		Name:      "stats_messages_sent_in_bytes",
		Help:      "Total number of bytes that cluster stats messages sent.",
		Labels:    []string{"addr"},
	},
	"cluster_stats_messages_ping_received": &MetricDesc{
		Subsystem: "cluster",
		Name:      "cluster_stats_messages_ping_received_in_bytes",
		Help:      "Total number of bytes that cluster stats messages ping received.",
		Labels:    []string{"addr"},
	},
	"cluster_stats_messages_pong_received": &MetricDesc{
		Subsystem: "cluster",
		Name:      "cluster_stats_messages_pong_received_in_bytes",
		Help:      "Total number of bytes that cluster stats messages pong received.",
		Labels:    []string{"addr"},
	},
	"cluster_stats_messages_publish_received": &MetricDesc{
		Subsystem: "cluster",
		Name:      "stats_messages_publish_received_in_bytes",
		Help:      "Total number of bytes that cluster stats messages publish received.",
		Labels:    []string{"addr"},
	},
	"cluster_stats_messages_received": &MetricDesc{
		Subsystem: "cluster",
		Name:      "stats_messages_received_in_bytes",
		Help:      "Total number of bytes that cluster stats messages received.",
		Labels:    []string{"addr"},
	},
}

type clusterInfoScraper struct {
	metricsDesc map[string]*MetricDesc
}

func NewClusterInfoScraper() *clusterInfoScraper {
	return &clusterInfoScraper{
		metricsDesc: clusterInfoMetricsDesc,
	}
}

// Scrape implements Scraper.
func (scraper *clusterInfoScraper) Scrape(ctx context.Context, rdbs []*redis.Client, ch chan<- prometheus.Metric, logger log.Logger) error {
	var err error

	for _, rdb := range rdbs {
		addr := rdb.Options().Addr

		var res string
		res, err = rdb.ClusterInfo(ctx).Result()
		if err != nil {
			return err
		}

		resMap := parseRedisRespInfo(res)

		for k, v := range scraper.metricsDesc {
			var f64 float64
			f64, err = strconv.ParseFloat(resMap[k], 64)
			checkParseRedisRespInfoError(k, addr, err, logger)

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

// Help implements Scraper.
func (*clusterInfoScraper) Help() string {
	return "Collect cluster info from each redis node."
}

// Name implements Scraper.
func (*clusterInfoScraper) Name() string {
	return "cluster.info"
}

// Version implements Scraper.
func (*clusterInfoScraper) Version() string {
	return "3.0"
}

var _ Scraper = &clusterInfoScraper{}
