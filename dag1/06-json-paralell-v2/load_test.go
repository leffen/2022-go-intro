package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {

	chMeasurement := make(chan *Measurement)

	go func() {
		err := LoadWeatherData("../data/alldata.json", chMeasurement)
		assert.Nil(t, err)
	}()

	rows := []*Measurement{}
	rowNum := 0
	for m := range chMeasurement {
		rows = append(rows, m)
		rowNum++
		if rowNum%10000 == 0 {
			fmt.Print("#")
		}
	}
	assert.Equal(t, 572922, len(rows))

}

func BenchmarkTestLoadData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chMeasurement := make(chan *Measurement)
		go func() {
			err := LoadWeatherData("../data/alldata.json", chMeasurement)
			assert.Nil(b, err)
		}()

		rows := []*Measurement{}
		rowNum := 0
		for m := range chMeasurement {
			rows = append(rows, m)
			rowNum++
		}
		assert.Equal(b, 572922, len(rows))
	}
}
