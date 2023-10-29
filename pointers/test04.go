package main

import "fmt"

func main() {
	a := 10

	var pointer *int = &a
	*pointer = 22

	fmt.Println(a)
}
