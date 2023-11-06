package main

import (
	ctx "context"
	"fmt"
	"time"
)

func main() {
	context, cancel := ctx.WithCancel(ctx.Background())
	defer cancel()

	go func() {
		time.Sleep(time.Second * 8) // Requisition time
		cancel()
	}()
	check(context)
}

func check(context ctx.Context) {
	select {
	case <-context.Done():
		fmt.Println("That is done!")

	case <-time.After(time.Second * 5):
		fmt.Println("Time out... Try again later.")
	}
}
