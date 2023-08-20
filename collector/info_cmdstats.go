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

import "fmt"

func initCmdStatsMetricsDesc(m map[string]string) map[string]*MetricDesc {
	var cmdStatsMetricsDesc map[string]*MetricDesc = make(map[string]*MetricDesc, 1)
	for k, _ := range m {
		cmdStatsMetricsDesc[k] = &MetricDesc{
			Subsystem: "server",
			Name:      k,
			Help:      fmt.Sprintf("The stats of %s.", k),
			Labels:    []string{"addr"},
		}
	}

	return cmdStatsMetricsDesc
}

func NewInfoCommandStatsScraper() *infoScraper {
	return &infoScraper{
		section:     "commandstats",
		sectionHelp: "Collect info commandstats from each redis server.",
	}
}
