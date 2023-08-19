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

var clientsMetricsDesc map[string]*MetricDesc = map[string]*MetricDesc{
	"connected_clients": &MetricDesc{
		Subsystem: "server",
		Name:      "connected_clients_in_total",
		Help:      "Number of client connections (excluding connections from replicas).",
		Labels:    []string{"addr"},
	},
	"blocked_clients": &MetricDesc{
		Subsystem: "server",
		Name:      "blocked_clients_in_total",
		Help:      "Number of clients pending on a blocking call (BLPOP, BRPOP, BRPOPLPUSH, BLMOVE, BZPOPMIN, BZPOPMAX).",
		Labels:    []string{"addr"},
	},
	"client_recent_max_input_buffer": &MetricDesc{
		Subsystem: "server",
		Name:      "client_recent_max_input_buffer_in_bytes",
		Help:      "Biggest input buffer among current client connections.",
		Labels:    []string{"addr"},
	},
	"client_recent_max_output_buffer": &MetricDesc{
		Subsystem: "server",
		Name:      "client_recent_max_output_buffer_in_bytes",
		Help:      "Biggest output buffer among current client connections.",
		Labels:    []string{"addr"},
	},
}

func NewInfoClientsScraper() *infoScraper {
	return &infoScraper{
		section:     "clients",
		sectionHelp: "Collect info clients from each redis server.",
		metricsDesc: clientsMetricsDesc,
	}
}
