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
