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

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))

	http.HandleFunc("/measurement", mhs.HandleRequest)
	http.HandleFunc("/hello", mhs.HandleIndex)
	http.HandleFunc("/", mhs.HandleIndex01)

	fmt.Printf("Startet server på port %s\n", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		logrus.Fatal(err)
	}
}
