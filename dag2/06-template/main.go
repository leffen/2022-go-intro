package main

import (
	"os"
	"text/template"

	"github.com/sirupsen/logrus"
)

func main() {
	tpl, err := template.ParseFiles("template-01.txt")
	if err != nil {
		logrus.Fatal(err)
	}

	data := map[string]string{"Name": "Test"}

	tpl.Execute(os.Stdout, data)
}
