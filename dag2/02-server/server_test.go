package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostEndpoint(t *testing.T) {

	jsonData, err := ioutil.ReadFile("testdata/data.json")
	require.Nil(t, err)
	require.NotNil(t, jsonData)

	s := NewMeasurementHttpServer()
	for i := 0; i < 100000; i++ {
		request, err := http.NewRequest("POST", "/measurement", strings.NewReader(string(jsonData)))
		require.Nil(t, err)
		require.NotNil(t, request)

		response := httptest.NewRecorder()

		s.HandleRequest(response, request)

		assert.Equal(t, 200, response.Code, "OK is expected when valid data is provided")

	}
	assert.Equal(t, 100000, len(s.items))

	//assert.True(t, false) // DEBUG
}
