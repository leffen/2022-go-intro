package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
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

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	p := &ElasticPublisher{}
	err := p.Connect("http://localhost:9200", "weatherdata")
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Connected to %s", p.url)

	chMeasurement := make(chan *Measurement)

	go func() {
		err := LoadWeatherData("../data/alldata.json", chMeasurement)
		if err != nil {
			logrus.Fatal(err)
		}
	}()

	rowNum := 0
	for m := range chMeasurement {
		js, err := json.Marshal(m)
		if err != nil {
			logrus.Error("Unable to marshal %s", err)
			continue
		}
		p.Publish("weather", string(js))
		rowNum++
		if rowNum%10000 == 0 {
			fmt.Print("#")
		}
	}

}

func LoadWeatherData(fileName string, ch chan *Measurement) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wg := sync.WaitGroup{}
	chData := make(chan string)

	numChannels := 100

	// Start x workers
	for i := 0; i < numChannels; i++ {
		go worker(chData, ch, &wg)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}
		wg.Add(1)
		chData <- line
	}
	wg.Wait()
	close(ch)

	return nil
}

func worker(chIn chan string, chOut chan *Measurement, wg *sync.WaitGroup) {
	for data := range chIn {
		m := &Measurement{}
		err := json.Unmarshal([]byte(data), m)
		if err != nil {
			logrus.Error(err)
		}
		chOut <- m
		wg.Done()
	}
}
