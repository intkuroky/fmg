package fmg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	InitConfig("")
	assert.Equal(t, DefaultConfigFile, Config.Viper.ConfigFileUsed())
}
