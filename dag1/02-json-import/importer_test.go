package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadFile(t *testing.T) {
	data, err := loadMessage("florida.json")
	require.Nil(t, err)
	assert.NotNil(t, data)
}
