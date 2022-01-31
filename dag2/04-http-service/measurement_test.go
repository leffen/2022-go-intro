package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMeasurement(t *testing.T) {
	jsonData, err := ioutil.ReadFile("testdata/data.json")
	require.Nil(t, err)
	require.NotNil(t, jsonData)

	measurement, err := NewMesurementFromJSON(jsonData)
	require.Nil(t, err)
	require.NotNil(t, measurement)
}
