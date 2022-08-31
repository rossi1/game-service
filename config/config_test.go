package config

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg, err := GetConfig("./..")
	assert.NoError(t, err)

	t.Run("test environment", func(t *testing.T) {
		if !cfg.Environment.isvalid() {
			t.Fatal("Invalid environment")
		}

	})

	t.Run("test dep", func(t *testing.T) {
		dep, err := GetDeps(context.Background(), cfg)
		assert.NoError(t, err)
		assert.NotNil(t, dep)

	})

}
