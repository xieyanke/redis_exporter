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

var replicationMetricsDesc = map[string]*MetricDesc{
	"connected_slaves": &MetricDesc{
		Subsystem: "server",
		Name:      "connected_slaves_in_total",
		Help:      "Number of connected replicas.",
		Labels:    []string{"addr"},
	},
	"repl_backlog_active": &MetricDesc{
		Subsystem: "server",
		Name:      "repl_backlog_active_status",
		Help:      "Flag indicating replication backlog is active.",
		Labels:    []string{"addr"},
	},
	"repl_backlog_size": &MetricDesc{
		Subsystem: "server",
		Name:      "repl_backlog_size_in_bytes",
		Help:      "Total size in bytes of the replication backlog buffer.",
		Labels:    []string{"addr"},
	},
	"repl_backlog_first_byte_offset": &MetricDesc{
		Subsystem: "server",
		Name:      "repl_backlog_first_byte_offset",
		Help:      "The master offset of the replication backlog buffer.",
		Labels:    []string{"addr"},
	},
	"repl_backlog_histlen": &MetricDesc{
		Subsystem: "server",
		Name:      "repl_backlog_histlen",
		Help:      "Size in bytes of the data in the replication backlog buffer.",
		Labels:    []string{"addr"},
	},
}

func NewInfoReplicationScraper() *InfoScraper {
	return &InfoScraper{
		section:     "replication",
		sectionHelp: "Collect info replication from each redis server.",
		metricsDesc: replicationMetricsDesc,
	}
}
