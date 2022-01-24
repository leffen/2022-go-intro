package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadData(t *testing.T) {
	data, err := LoadWeatherData("testdata.json")

	require.Nil(t, err)

	assert.Equal(t, 10, len(data))

	e1 := data[0]
	assert.Equal(t, "EasyWeatherV1.5.4", e1.StationType)

}
