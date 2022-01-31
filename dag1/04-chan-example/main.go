package main

import "fmt"

func main() {
	chMessages := make(chan string,10)

	go lagMelding(chMessages)

	for msg := range chMessages {
		fmt.Printf("Received: %s\n",msg)
	}
	fmt.Println("Done")
}

func lagMelding(chMessages chan string) {
	i:=0
	for i<10 {
		chMessages <- fmt.Sprintf("Message %d",i)
		fmt.Printf("sending message %d\n",i)
		i++
	}
	close(chMessages)
}
