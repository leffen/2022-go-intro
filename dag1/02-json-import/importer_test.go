package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadFile(t *testing.T) {
	err := loadMessage("test1.json")
	require.Nil(t, err)
}
