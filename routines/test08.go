package main

import (
	"fmt"
	"time"
)

func worker(workerID int, message chan int) {
	for response := range message {
		fmt.Println("Worker:", workerID, "working on number:", response)
		time.Sleep(time.Second)
	}
}

// Thread 1
func main() {
	number := make(chan int)

	// Thread 2
	go worker(1, number)
	// Thread 3
	go worker(2, number)

	for i := 0; i < 10; i++ {
		number <- i
	}

	// Only so we can see
	time.Sleep(time.Second * 5)
}
