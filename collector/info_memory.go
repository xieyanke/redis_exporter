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

var memoryMetricsDesc = map[string]*MetricDesc{
	"used_memory": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_in_bytes",
		Help:      "Total number of bytes allocated by Redis using its allocator (either standard libc, jemalloc, or an alternative allocator such as tcmalloc).",
		Labels:    []string{"addr"},
	},
	"used_memory_rss": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_rss_in_bytes",
		Help:      "Number of bytes that Redis allocated as seen by the operating system (a.k.a resident set size). This is the number reported by tools such as top(1) and ps(1).",
		Labels:    []string{"addr"},
	},
	"used_memory_peak": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_peak_in_bytes",
		Help:      "Peak memory consumed by Redis (in bytes).",
		Labels:    []string{"addr"},
	},
	"used_memory_overhead": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_overhead_in_bytes",
		Help:      "The sum in bytes of all overheads that the server allocated for managing its internal data structures.",
		Labels:    []string{"addr"},
	},
	"used_memory_startup": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_startup_in_bytes",
		Help:      "Initial amount of memory consumed by Redis at startup in bytes.",
		Labels:    []string{"addr"},
	},
	"used_memory_dataset": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_dataset_in_bytes",
		Help:      "The size in bytes of the dataset (used_memory_overhead subtracted from used_memory).",
		Labels:    []string{"addr"},
	},
	"allocator_allocated": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_allocator_allocated_in_bytes",
		Help:      "Total bytes allocated form the allocator, including internal-fragmentation. Normally the same as used_memory.",
		Labels:    []string{"addr"},
	},
	"allocator_active": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_allocator_active_in_bytes",
		Help:      "Total bytes in the allocator active pages, this includes external-fragmentation.",
		Labels:    []string{"addr"},
	},
	"allocator_resident": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_allocator_resident_in_bytes",
		Help:      "Total bytes resident (RSS) in the allocator, this includes pages that can be released to the OS (by MEMORY PURGE, or just waiting).",
		Labels:    []string{"addr"},
	},
	"total_system_memory": &MetricDesc{
		Subsystem: "server",
		Name:      "host_total_system_memory_in_bytes",
		Help:      "The total amount of memory that the Redis host has.",
		Labels:    []string{"addr"},
	},
	"used_memory_lua": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_lua_in_bytes",
		Help:      "Number of bytes used by the Lua engine.",
		Labels:    []string{"addr"},
	},
	"used_memory_scripts": &MetricDesc{
		Subsystem: "server",
		Name:      "used_memory_scripts_in_bytes",
		Help:      "Number of bytes used by cached Lua scripts.",
		Labels:    []string{"addr"},
	},
	"number_of_cached_scripts": &MetricDesc{
		Subsystem: "server",
		Name:      "number_of_cached_scripts_in_total",
		Help:      "Number of cached Lua scripts.",
		Labels:    []string{"addr"},
	},
	"maxmemory": &MetricDesc{
		Subsystem: "server",
		Name:      "maxmemory_in_bytes",
		Help:      "The value of the maxmemory configuration directive.",
		Labels:    []string{"addr"},
	},
	"allocator_frag_ratio": &MetricDesc{
		Subsystem: "server",
		Name:      "allocator_frag_ratio",
		Help:      "Ratio between allocator_active and allocator_allocated. This is the true (external) fragmentation metric (not mem_fragmentation_ratio).",
		Labels:    []string{"addr"},
	},
	"allocator_frag_bytes": &MetricDesc{
		Subsystem: "server",
		Name:      "allocator_frag_in_bytes",
		Help:      "Delta between allocator_active and allocator_allocated. See note about mem_fragmentation_bytes.",
		Labels:    []string{"addr"},
	},
	"allocator_rss_ratio": &MetricDesc{
		Subsystem: "server",
		Name:      "allocator_rss_ratio",
		Help:      "Ratio between allocator_resident and allocator_active. This usually indicates pages that the allocator can and probably will soon release back to the OS.",
		Labels:    []string{"addr"},
	},
	"allocator_rss_bytes": &MetricDesc{
		Subsystem: "server",
		Name:      "allocator_rss_in_bytes",
		Help:      "Delta between allocator_resident and allocator_active.",
		Labels:    []string{"addr"},
	},
	"rss_overhead_ratio": &MetricDesc{
		Subsystem: "server",
		Name:      "rss_overhead_ratio",
		Help:      "Collect redis rss overhead ratio(the ratio of used_memory_rss and allocator_resident) from each redis server.",
		Labels:    []string{"addr"},
	},
	"rss_overhead_bytes": &MetricDesc{
		Subsystem: "server",
		Name:      "rss_overhead_in_bytes",
		Help:      "Ratio between used_memory_rss (the process RSS) and allocator_resident. This includes RSS overheads that are not allocator or heap related.",
		Labels:    []string{"addr"},
	},
	"mem_fragmentation_ratio": &MetricDesc{
		Subsystem: "server",
		Name:      "mem_fragmentation_ratio",
		Help:      "Ratio between used_memory_rss and used_memory. Note that this doesn't only includes fragmentation, but also other process overheads (see the allocator_* metrics), and also overheads like code, shared libraries, stack, etc.",
		Labels:    []string{"addr"},
	},
	"mem_fragmentation_bytes": &MetricDesc{
		Subsystem: "server",
		Name:      "mem_fragmentation_in_bytes",
		Help:      "Delta between used_memory_rss and used_memory. Note that when the total fragmentation bytes is low (few megabytes), a high ratio (e.g. 1.5 and above) is not an indication of an issue.",
		Labels:    []string{"addr"},
	},
	"mem_not_counted_for_evict": &MetricDesc{
		Subsystem: "server",
		Name:      "mem_not_counted_for_evict_in_bytes",
		Help:      "Used memory that's not counted for key eviction. This is basically transient replica and AOF buffers.",
		Labels:    []string{"addr"},
	},
	"mem_replication_backlog": &MetricDesc{
		Subsystem: "server",
		Name:      "mem_replication_backlog_in_bytes",
		Help:      "Memory used by replication backlog.",
		Labels:    []string{"addr"},
	},
	"mem_clients_slaves": &MetricDesc{
		Subsystem: "server",
		Name:      "mem_clients_slaves_in_bytes",
		Help:      "Memory used by replica clients - Starting Redis 7.0, replica buffers share memory with the replication backlog, so this field can show 0 when replicas don't trigger an increase of memory usage.",
		Labels:    []string{"addr"},
	},
	"mem_clients_normal": &MetricDesc{
		Subsystem: "server",
		Name:      "mem_clients_normal_in_bytes",
		Help:      "Memory used by normal clients.",
		Labels:    []string{"addr"},
	},
	"mem_aof_buffer": &MetricDesc{
		Subsystem: "server",
		Name:      "mem_aof_buffer_in_bytes",
		Help:      "Transient memory used for AOF and AOF rewrite buffers.",
		Labels:    []string{"addr"},
	},

	"active_defrag_running": &MetricDesc{
		Subsystem: "server",
		Name:      "active_defrag_running_in_total",
		Help:      "When activedefrag is enabled, this indicates whether defragmentation is currently active, and the CPU percentage it intends to utilize.",
		Labels:    []string{"addr"},
	},
	"lazyfree_pending_objects": &MetricDesc{
		Subsystem: "server",
		Name:      "lazyfree_pending_objects_in_total",
		Help:      "The number of objects waiting to be freed (as a result of calling UNLINK, or FLUSHDB and FLUSHALL with the ASYNC option).",
		Labels:    []string{"addr"},
	},
}

func NewInfoMemoryScraper() *infoScraper {
	return &infoScraper{
		section:     "memory",
		sectionHelp: "Collect info memory from each redis server.",
		metricsDesc: memoryMetricsDesc,
	}
}
