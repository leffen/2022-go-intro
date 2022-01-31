package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	items := []string{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			items = append(items, fmt.Sprintf("Item %d", i))
			mu.Unlock()
			time.Sleep(20 * time.Nanosecond)
		}(i)
	}
	wg.Wait()
	for item := range items {
		fmt.Println(item)
	}
}
