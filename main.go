package main

import "fmt"

func main() {
	a := []int{10, 10, 20}
	fmt.Println(sumArray(a))
}

func sumArray(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total
}
