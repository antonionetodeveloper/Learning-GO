package main

import (
	"fmt"
	"time"
)

func counter(word string, times int) {
	for i := 1; i <= times; i++ {
		fmt.Println(word, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go counter("Contador 1:", 5)
	go counter("Contador 2:", 5)
	time.Sleep(time.Second * 5)
}
