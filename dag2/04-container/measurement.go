package main

import (
	"encoding/json"
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

func NewMesurementFromJSON(data []byte) (*Measurement, error) {
	m := &Measurement{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
