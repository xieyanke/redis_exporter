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

var statsMetricsDesc = map[string]*MetricDesc{
	"total_connections_received": &MetricDesc{
		Subsystem: "server",
		Name:      "connections_received_in_total",
		Help:      "Total number of connections accepted by the server.",
		Labels:    []string{"addr"},
	},
	"total_commands_processed": &MetricDesc{
		Subsystem: "server",
		Name:      "commands_processed_in_total",
		Help:      "Total number of commands processed by the server.",
		Labels:    []string{"addr"},
	},

	"instantaneous_ops_per_sec": &MetricDesc{
		Subsystem: "server",
		Name:      "commands_processed_per_second",
		Help:      "Number of commands processed per second.",
		Labels:    []string{"addr"},
	},
	"total_net_input_bytes": &MetricDesc{
		Subsystem: "server",
		Name:      "total_net_input_bytes",
		Help:      "The total number of bytes read from the network.",
		Labels:    []string{"addr"},
	},
	"total_net_output_bytes": &MetricDesc{
		Subsystem: "server",
		Name:      "total_net_output_bytes",
		Help:      "The total number of bytes written to the network.",
		Labels:    []string{"addr"},
	},
	"instantaneous_input_kbps": &MetricDesc{
		Subsystem: "server",
		Name:      "instantaneous_input_kbps",
		Help:      "The network's read rate per second in KB/sec.",
		Labels:    []string{"addr"},
	},
	"instantaneous_output_kbps": &MetricDesc{
		Subsystem: "server",
		Name:      "instantaneous_output_kbps",
		Help:      "The network's write rate per second in KB/sec.",
		Labels:    []string{"addr"},
	},
	"rejected_connections": &MetricDesc{
		Subsystem: "server",
		Name:      "rejected_connections_in_total",
		Help:      "Number of connections rejected because of maxclients limit.",
		Labels:    []string{"addr"},
	},
	"sync_full": &MetricDesc{
		Subsystem: "server",
		Name:      "sync_full_in_total",
		Help:      "The number of full resyncs with replicas.",
		Labels:    []string{"addr"},
	},
	"sync_partial_ok": &MetricDesc{
		Subsystem: "server",
		Name:      "sync_partial_ok_in_total",
		Help:      "The number of accepted partial resync requests.",
		Labels:    []string{"addr"},
	},
	"sync_partial_err": &MetricDesc{
		Subsystem: "server",
		Name:      "sync_partial_err_in_total",
		Help:      "The number of denied partial resync requests.",
		Labels:    []string{"addr"},
	},
	"expired_keys": &MetricDesc{
		Subsystem: "server",
		Name:      "expired_keys_in_total",
		Help:      "Total number of key expiration events.",
		Labels:    []string{"addr"},
	},
	"expired_stale_perc": &MetricDesc{
		Subsystem: "server",
		Name:      "expired_stale_percent",
		Help:      "The percentage of keys probably expired.",
		Labels:    []string{"addr"},
	},
	"expired_time_cap_reached_count": &MetricDesc{
		Subsystem: "server",
		Name:      "expired_time_cap_reached_count",
		Help:      "The count of times that active expiry cycles have stopped early.",
		Labels:    []string{"addr"},
	},
	"evicted_keys": &MetricDesc{
		Subsystem: "server",
		Name:      "evicted_keys_in_total",
		Help:      "Number of evicted keys due to maxmemory limit.",
		Labels:    []string{"addr"},
	},
	"keyspace_hits": &MetricDesc{
		Subsystem: "server",
		Name:      "keyspace_hits_in_total",
		Help:      "Number of successful lookup of keys in the main dictionary.",
		Labels:    []string{"addr"},
	},
	"keyspace_misses": &MetricDesc{
		Subsystem: "server",
		Name:      "keyspace_misses_in_total",
		Help:      "Number of failed lookup of keys in the main dictionary.",
		Labels:    []string{"addr"},
	},
	"pubsub_channels": &MetricDesc{
		Subsystem: "server",
		Name:      "pubsub_channels_in_total",
		Help:      "Global number of pub/sub channels with client subscriptions.",
		Labels:    []string{"addr"},
	},
	"pubsub_patterns": &MetricDesc{
		Subsystem: "server",
		Name:      "pubsub_patterns_in_total",
		Help:      "Global number of pub/sub pattern with client subscriptions.",
		Labels:    []string{"addr"},
	},
	"latest_fork_usec": &MetricDesc{
		Subsystem: "server",
		Name:      "latest_fork_in_microseconds",
		Help:      "Duration of the latest fork operation in microseconds.",
		Labels:    []string{"addr"},
	},
	"migrate_cached_sockets": &MetricDesc{
		Subsystem: "server",
		Name:      "migrate_cached_sockets_in_total",
		Help:      "The number of sockets open for migrage purposes",
		Labels:    []string{"addr"},
	},
	"slave_expires_tracked_keys": &MetricDesc{
		Subsystem: "server",
		Name:      "slave_expires_tracked_keys_in_total",
		Help:      "The number of keys tracked for expiry purposes (applicable only to writable replicas).",
		Labels:    []string{"addr"},
	},
	"active_defrag_hits": &MetricDesc{
		Subsystem: "server",
		Name:      "active_defrag_hits_in_total",
		Help:      "Number of value reallocations performed by active the defragmentation process.",
		Labels:    []string{"addr"},
	},
	"active_defrag_misses": &MetricDesc{
		Subsystem: "server",
		Name:      "active_defrag_misses_in_total",
		Help:      "Number of aborted value reallocations started by the active defragmentation process.",
		Labels:    []string{"addr"},
	},
	"active_defrag_key_hits": &MetricDesc{
		Subsystem: "server",
		Name:      "active_defrag_key_hits_in_total",
		Help:      "Number of keys that were actively defragmented.",
		Labels:    []string{"addr"},
	},
	"active_defrag_key_misses": &MetricDesc{
		Subsystem: "server",
		Name:      "active_defrag_key_misses_in_total",
		Help:      "Number of keys that were skipped by the active defragmentation process.",
		Labels:    []string{"addr"},
	},
}

func NewInfoStatsScraper() *infoScraper {
	return &infoScraper{
		section:     "stats",
		sectionHelp: "Collect info stats from each redis server.",
		metricsDesc: statsMetricsDesc,
	}
}
