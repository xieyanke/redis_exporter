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

var persistenceMetricsDesc = map[string]*MetricDesc{
	"loading": &MetricDesc{
		Subsystem: "server",
		Name:      "persistence_loading_status",
		Help:      "Flag indicating if the load of a dump file is on-going.(1: yes, 0: no)",
		Labels:    []string{"addr"},
	},
	"rdb_changes_since_last_save": &MetricDesc{
		Subsystem: "server",
		Name:      "rdb_changes_since_last_save_in_total",
		Help:      "Number of changes since the last dump.",
		Labels:    []string{"addr"},
	},
	"rdb_last_save_time": &MetricDesc{
		Subsystem: "server",
		Name:      "rdb_last_save_timetamp",
		Help:      "Epoch-based timestamp of last successful RDB save.",
		Labels:    []string{"addr"},
	},
	"rdb_last_bgsave_status": &MetricDesc{
		Subsystem: "server",
		Name:      "rdb_last_bgsave_status",
		Help:      "Status of the last RDB save operation.",
		Labels:    []string{"addr"},
	},
	"rdb_last_bgsave_time_sec": &MetricDesc{
		Subsystem: "server",
		Name:      "rdb_last_bgsave_used_seconds",
		Help:      "Duration of the last RDB save operation in seconds.",
		Labels:    []string{"addr"},
	},
	"rdb_current_bgsave_time_sec": &MetricDesc{
		Subsystem: "server",
		Name:      "rdb_current_bgsave_used_seconds",
		Help:      "Duration of the on-going RDB save operation if any.",
		Labels:    []string{"addr"},
	},

	"rdb_last_cow_size": &MetricDesc{
		Subsystem: "server",
		Name:      "rdb_last_cow_size_in_bytes",
		Help:      "The size in bytes of copy-on-write memory during the last RDB save operation",
		Labels:    []string{"addr"},
	},
	"aof_enabled": &MetricDesc{
		Subsystem: "server",
		Name:      "aof_enabled",
		Help:      "Flag indicating AOF logging is activated.",
		Labels:    []string{"addr"},
	},
	"aof_rewrite_in_progress": &MetricDesc{
		Subsystem: "server",
		Name:      "aof_rewrite_in_progress_status",
		Help:      "Flag indicating a AOF rewrite operation is on-going.",
		Labels:    []string{"addr"},
	},
	"aof_rewrite_scheduled": &MetricDesc{
		Subsystem: "server",
		Name:      "aof_rewrite_scheduled_status",
		Help:      "Flag indicating an AOF rewrite operation will be scheduled once the on-going RDB save is complete.",
		Labels:    []string{"addr"},
	},
	"aof_last_rewrite_time_sec": &MetricDesc{
		Subsystem: "server",
		Name:      "aof_last_rewrite_time_in_seconds",
		Help:      "Duration of the last AOF rewrite operation in seconds.",
		Labels:    []string{"addr"},
	},
	"aof_current_rewrite_time_sec": &MetricDesc{
		Subsystem: "server",
		Name:      "aof_current_rewrite_time_in_seconds",
		Help:      "Duration of the on-going AOF rewrite operation if any.",
		Labels:    []string{"addr"},
	},
	"aof_last_bgrewrite_status": &MetricDesc{
		Subsystem: "server",
		Name:      "aof_last_bgrewrite_status",
		Help:      "Status of the last AOF rewrite operation.",
		Labels:    []string{"addr"},
	},
	"aof_last_write_status": &MetricDesc{
		Subsystem: "server",
		Name:      "aof_last_write_status",
		Help:      "Status of the last write operation to the AOF.",
		Labels:    []string{"addr"},
	},
	"aof_last_cow_size": &MetricDesc{
		Subsystem: "server",
		Name:      "aof_last_cow_size_in_bytes",
		Help:      "The size in bytes of copy-on-write memory during the last AOF rewrite operation.",
		Labels:    []string{"addr"},
	},
}

func NewInfoPersistenceScraper() *InfoScraper {
	return &InfoScraper{
		section:     "persistence",
		sectionHelp: "Collect info persistence from each redis server.",
		metricsDesc: persistenceMetricsDesc,
	}
}
