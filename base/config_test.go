package base

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseConfigWithPrefix(t *testing.T) {
	type mockStruct struct {
		A string `envconfig:"A"`
		B int    `envconfig:"B"`
	}
	assert.NoError(t, os.Setenv("MOCK_A", "a"))
	assert.NoError(t, os.Setenv("MOCK_B", "1"))
	cfg, err := ParseConfigWithPrefix[mockStruct]("MOCK")
	assert.NoError(t, err)
	assert.Equal(t, "a", cfg.A)
	assert.Equal(t, 1, cfg.B)
}

func TestParseConfig(t *testing.T) {
	type mockStruct struct {
		A string `envconfig:"A"`
		B int    `envconfig:"B"`
	}
	assert.NoError(t, os.Setenv("A", "a"))
	assert.NoError(t, os.Setenv("B", "1"))
	cfg, err := ParseConfig[mockStruct]()
	assert.NoError(t, err)
	assert.Equal(t, "a", cfg.A)
	assert.Equal(t, 1, cfg.B)
}
