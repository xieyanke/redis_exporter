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

var serverMetricsDesc map[string]*MetricDesc = map[string]*MetricDesc{
	"uptime_in_seconds": &MetricDesc{
		Subsystem: "server",
		Name:      "uptime_in_seconds",
		Help:      "Number of seconds since Redis server start.",
		Labels:    []string{"addr"},
	},

	"hz": &MetricDesc{
		Subsystem: "server",
		Name:      "hz",
		Help:      "The server's current frequency setting.",
		Labels:    []string{"addr"},
	},
	"configured_hz": &MetricDesc{
		Subsystem: "server",
		Name:      "configured_hz",
		Help:      "The server's configured frequency setting.",
		Labels:    []string{"addr"},
	},
	"lru_clock": &MetricDesc{
		Subsystem: "server",
		Name:      "lru_clock",
		Help:      "Clock incrementing every minute, for LRU management.",
		Labels:    []string{"addr"},
	},
}

func NewInfoServerScraper() *infoScraper {
	return &infoScraper{
		section:     "server",
		sectionHelp: "Collect info server from each redis server.",
		metricsDesc: serverMetricsDesc,
	}
}
