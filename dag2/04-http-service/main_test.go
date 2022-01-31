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

	request, err := http.NewRequest("POST", "/measurement", strings.NewReader(string(jsonData)))
	require.Nil(t, err)
	require.NotNil(t, request)

	response := httptest.NewRecorder()

	handlePost(response, request)

	assert.Equal(t, 200, response.Code, "OK is expected when valid data is provided")
}

func TestPostNilToEndpoint(t *testing.T) {

	request, err := http.NewRequest("POST", "/measurement", nil)
	require.Nil(t, err)

	response := httptest.NewRecorder()

	handlePost(response, request)
	assert.Equal(t, 400, response.Code, "Bad request expected when no data is provided")
}
