package util

import (
	"strings"
)

func ParseInfoSection(section string) map[string]string {
	section = strings.TrimSpace(section)
	lines := strings.Split(section, "\n")

	m := make(map[string]string, 4)
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			arr := strings.Split(line, ":")
			m[arr[0]] = arr[1]
		}
	}

	return m
}
