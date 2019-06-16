package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitConfig(t *testing.T) {
	tests := []struct{
		configName string
		expect error
	}{
		{"", nil},
		{"config", nil},
	}
	for _, test := range tests {
		assert.Equal(t, initConfig(test.configName), test.expect)
	}

	configName := "test"
	assert.NotNil(t, initConfig(configName))
}