package main

import "fmt"

// Thread 1
func main() {
	// Thread 1 <-> Thread 2
	hello := make(chan string)

	// Thread 2
	go func() {
		hello <- "Hello from other channel."
	}()

	result := <-hello
	fmt.Println(result)
}
