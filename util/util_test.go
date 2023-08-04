package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data string = `
# Server
redis_version:5.0.14
redis_git_sha1:00000000
`

func TestParseInfoSection(t *testing.T) {
	expected := map[string]string{
		"redis_version":  "5.0.14",
		"redis_git_sha1": "00000000",
	}

	actual := ParseInfoSection(data)
	assert.Equal(t, actual, expected)
}
