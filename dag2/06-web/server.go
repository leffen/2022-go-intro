package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
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

func (s *MeasurementHttpServer) HandleIndex(w http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseFiles("template/base.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	list := []string{}
	list = append(list, "<ul>\n")

	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, m := range s.items {
		list = append(list, fmt.Sprintf("<li>%v - vind: %v</li>\n", m.Date, m.WindSpeedMS))
	}
	list = append(list, "</ul>\n")

	w.Header().Set("Content-Type", "text/html")
	tpl.Execute(w, template.HTML(strings.Join(list, "\n")))
}

func (s *MeasurementHttpServer) HandleRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		s.Post(w, req)
		return
	}
	if req.Method == "GET" {
		s.Get(w, req)
		return
	}

	http.Error(w, "Unsupported", http.StatusBadRequest)
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

func (s *MeasurementHttpServer) Get(w http.ResponseWriter, req *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	jsonData, err := json.Marshal(s.items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(jsonData))

}
