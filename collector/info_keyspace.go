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
	"fmt"
	"strings"
)

func initKeyspaceMetricsDesc(m map[string]string) map[string]*MetricDesc {
	var keyspaceMetricsDesc map[string]*MetricDesc = make(map[string]*MetricDesc, 1)
	for k, _ := range m {
		var name string
		var help string

		if strings.Contains(k, "keys") {
			name = fmt.Sprintf("%s_%s_%s", "keyspace", k, "in_total")
			help = "Number of keyspace keys in the redis db"
		} else if strings.Contains(k, "expires") {
			name = fmt.Sprintf("%s_%s_%s", "keyspace", k, "in_total")
			help = "Number of expire keys in the redis db"
		} else if strings.Contains(k, "avg_ttl") {
			name = fmt.Sprintf("%s_%s_%s", "keyspace", k, "in_microseconds")
			help = "The average microseconds to live of all keys in the redis db."
		}

		keyspaceMetricsDesc[k] = &MetricDesc{
			Subsystem: "server",
			Name:      name,
			Help:      help,
			Labels:    []string{"addr"},
		}
	}

	return keyspaceMetricsDesc
}

func NewInfoKeyspaceScraper() *infoScraper {
	return &infoScraper{
		section:     "keyspace",
		sectionHelp: "Collect info keyspace from each redis server.",
	}
}
