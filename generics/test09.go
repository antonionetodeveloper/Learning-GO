package main

import "fmt"

func GenericSum[Type int64 | float64](array []Type) (result Type) {
	for _, value := range array {
		result += value
	}
	return
}

func main() {
	fmt.Println(GenericSum([]int64{1, 2, 3, 4, 5}))
	fmt.Println(GenericSum([]float64{1.3, 2.7, 3.1, 4.1, 5.2}))
}
