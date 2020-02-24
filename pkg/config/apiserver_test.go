package config

import (
	"gotest.tools/assert"
	"testing"
)

func TestLoadServer(t *testing.T) {
	LoadServer()
	assert.Equal(t, "http://localhost:8000/qurl/", Server().Qurl.Prefix)
	assert.Equal(t, 500, Server().Qurl.CacheSize)
}
