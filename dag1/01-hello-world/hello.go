package main

import (
	"fmt"
	"os"
)

func main() {
	var num int = 100
	var name string = "Ole Brum"

	// Utskrift av linje med parametre
	fmt.Printf("Hello %s %d\n", name, num)

	fmt.Println("Test av utskrift av linjer")

	// Utskrift til stderr
	fmt.Fprintf(os.Stderr, "Test av utskrift til stderr. Name=%s\n", name)

}
