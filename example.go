package main

import (
	"fmt"
	"time"
)

func demonstrateMapPanic() {
	m := make(map[int]bool)

	for i := 0; i < 100; i++ {
		go func() {
			m[i] = true
		}()
	}
}

func demonstrateRaceCondition() {
	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Expected: 1000, Got: %d\n", counter)
}

func main() {
	//demonstrateMapPanic()
	demonstrateRaceCondition()
}
