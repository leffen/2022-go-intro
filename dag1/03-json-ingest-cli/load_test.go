package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadFile(t *testing.T) {
	data, err := LoadWeatherData("test2.json")
	require.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, 100, len(data))
}
