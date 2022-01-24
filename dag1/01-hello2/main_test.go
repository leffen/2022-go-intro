package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJson(t *testing.T) {
	test := Test{Name: "TEST"}

	fmt.Printf("Data: %#v\n", test)

	require.NotNil(t, test, "Må være satt")

	jdata, err := test.ToJson()
	assert.Nil(t, err)
	assert.Equal(t, "{\"Name\":\"TEST\",\"Alder\":0}", jdata)

	//assert.True(t, false)
}
