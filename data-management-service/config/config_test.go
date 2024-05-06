package config_test

import (
	"data-management/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("Success Load Config", func(t *testing.T) {
		err := config.LoadConfig("../")
		assert.Nil(t, err)
	})

	t.Run("Fail Load Config", func(t *testing.T) {
		err := config.LoadConfig("invalid_path")
		assert.NotNil(t, err)
	})

	t.Run("Success Get Config", func(t *testing.T) {
		err := config.LoadConfig("../")
		assert.Nil(t, err)

		conf := config.GetConfig()
		assert.NotNil(t, conf)
	})

	t.Run("Not load config before get return nil", func(t *testing.T) {
		config.SetConfig(nil)
		conf := config.GetConfig()
		assert.Nil(t, conf)
	})

}
