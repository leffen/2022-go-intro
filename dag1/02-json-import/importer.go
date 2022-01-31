package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Message struct {
	Version  string  `json:"version"`
	Host     string  `json:"host"`
	Short    string  `json:"short_message"`
	Full     string  `json:"full_message,omitempty"`
	TimeUnix float64 `json:"timestamp"`
	Level    int32   `json:"level,omitempty"`
	Facility string  `json:"facility,omitempty"`
	Test     string  `json:"-"`
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Missing filename")
	}

	err := loadMessage(args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message successfully loaded")

}

func loadMessage(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	msg := Message{}
	err = json.Unmarshal(data, &msg)
	if err != nil {
		return err
	}

	fmt.Printf("Data: %+v\n", msg)
	return nil
}
