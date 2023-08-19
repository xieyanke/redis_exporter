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

var cpuMetricsDesc map[string]*MetricDesc = map[string]*MetricDesc{
	"used_cpu_sys": &MetricDesc{
		Subsystem: "server",
		Name:      "sys_cpu_used_in_total",
		Help:      "System CPU consumed by the Redis server, which is the sum of system CPU consumed by all threads of the server process (main thread and background threads).",
		Labels:    []string{"addr"},
	},
	"used_cpu_user": &MetricDesc{
		Subsystem: "server",
		Name:      "user_cpu_used_in_total",
		Help:      "User CPU consumed by the Redis server, which is the sum of user CPU consumed by all threads of the server process (main thread and background threads).",
		Labels:    []string{"addr"},
	},
	"used_cpu_sys_children": &MetricDesc{
		Subsystem: "server",
		Name:      "used_cpu_sys_children_in_total",
		Help:      "System CPU consumed by the background processes.",
		Labels:    []string{"addr"},
	},
	"used_cpu_user_children": &MetricDesc{
		Subsystem: "server",
		Name:      "used_cpu_user_children_in_total",
		Help:      "User CPU consumed by the background processes.",
		Labels:    []string{"addr"},
	},
}

func NewInfoCPUScraper() *infoScraper {
	return &infoScraper{
		section:     "cpu",
		sectionHelp: "Collect info cpu from each redis server.",
		metricsDesc: cpuMetricsDesc,
	}
}
