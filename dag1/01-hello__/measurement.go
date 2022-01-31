package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
	"time"
)

type Measurement struct {
	Date           time.Time
	Location       string  // From where is the measurement aquired
	StationType    string  // =EasyWeatherV1.5.3
	Dateutc        string  //
	TempIndoorC    float64 // Tempeature indoors
	HumidityIN     int64   // Humidity indoors
	HumidityAbs    float64 // Abselute hummidity
	Pressurerel    float64 // Pressure indoors
	Pressureabs    float64 // Absolute Pressure indoors
	TempOutdoorC   float64 // Temperature outdoor in C
	Humidity       int64   // Humidity outdoors
	WindDir        int64   // Wind directions
	WindSpeedMS    float64 // Windspeed in MS
	WindGustMS     float64 // WindGust in MS ( For what duration ?? )
	MaxDailyGustMS float64 // Maximum daily guest
	RainRate       float64 // RainRateMM
	EventRain      float64 // EventRain ?? in MM
	HourlyRain     float64 // Rain last hour in MM
	DailyRain      float64 // Daily rain in MM
	WeeklyRain     float64 // Weekly rain in MM
	MonthlyRain    float64 // Monthly rain in MM
	TotalRain      float64 // Total rain in MM
	SolarRadiation float64 // =79.76
	UV             int64   // =0
}

func LoadWeatherData(fileName string) ([]*Measurement, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rc := []*Measurement{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}

		m := &Measurement{}
		err := json.Unmarshal([]byte(line), m)
		if err != nil {
			return nil, err
		}
		rc = append(rc, m)
	}

	return rc, nil
}
