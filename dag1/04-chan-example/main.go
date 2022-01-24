package main

import "fmt"

func main() {
	chMessages := make(chan string)

	go lagMelding(chMessages)

	msg := <-chMessages

	fmt.Println(msg)

}

func lagMelding(chMessages chan string) {
	chMessages <- "ping"
}
