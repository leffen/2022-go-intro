package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	srv := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/measurement", handlePost)

	fmt.Printf("Startet server på port %s\n", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		logrus.Fatal(err)
	}
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
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

		fmt.Printf("Measurement: %#v\n", measurement)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
		return
	}

	http.Error(w, "Unhandled request", http.StatusBadRequest)
}
