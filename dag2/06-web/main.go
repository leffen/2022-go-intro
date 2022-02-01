package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	srv := &http.Server{
		Addr: ":8080",
	}

	mhs := NewMeasurementHttpServer()

	http.HandleFunc("/measurement", mhs.HandleRequest)
	http.HandleFunc("/", mhs.HandleIndex)

	fmt.Printf("Startet server på port %s\n", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		logrus.Fatal(err)
	}
}
