package main

import "fmt"

func main() {
	chMessages := make(chan string)
	go func() {
		chMessages <- "ping"
	}()
	msg := <-chMessages

	fmt.Println(msg)

}
