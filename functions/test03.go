package main

func main() {
	calculated := SumEverything(4, 7, 9, 12, 11)
	println(calculated)
}

func SumEverything(list ...int) (result int) {
	for _, value := range list {
		result += value
	}

	return
}
