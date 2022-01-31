package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type MeasurementHttpServer struct {
	items []*Measurement
	mu    sync.RWMutex
}

func NewMeasurementHttpServer() *MeasurementHttpServer {
	return &MeasurementHttpServer{
		items: []*Measurement{},
	}
}

func (s *MeasurementHttpServer) HandleRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		s.Post(w, req)
		return
	}

}

func (s *MeasurementHttpServer) Post(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "Missing body data", http.StatusBadRequest)
		return
	}

	jsonData, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	measurement, err := NewMesurementFromJSON(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.items = append(s.items, measurement)
	s.mu.Unlock()

	fmt.Printf("Measurement added\n")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
