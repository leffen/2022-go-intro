package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Test struct {
	Name  string
	Alder int
}

func (t *Test) Hello() {
	t.Alder = 4
	fmt.Printf("hello %d\n", t.Alder)
}

func (t *Test) SjekkAlder() error {
	if t.Alder < 18 {
		return fmt.Errorf("For ung. Må være mins 18 år")
	}
	return nil
}

func (t *Test) Kategori() (string, error) {
	if t.Alder < 4 {
		return "", fmt.Errorf("Ikke gyldig kategori")
	}

	if t.Alder < 19 {
		return "Tenåring", nil
	}

	return "Voksen", nil
}

func (t *Test) ToJson() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func NewFromJSON(data []byte) (*Test, error) {
	rc := Test{}
	err := json.Unmarshal(data, &rc)
	if err != nil {
		return nil, err
	}
	return &rc, nil
}

func xmain() {
	// obj := Test{Name: "Test"}
	// obj.Name = "Test21"
	// obj.Hello()

	// fmt.Printf("Data:%#v\n", obj)

	// jdata, err := obj.ToJson()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("JSON:", jdata)

	// o2, err := NewFromJSON([]byte(jdata))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("O2: %#v\n", o2)

	data, err := LoadWeatherData("../data/alldata.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Antall rader: %d\n", len(data))

}
