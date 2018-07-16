package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)
	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println("even", num)
		} else {
			fmt.Println("odd", num)
		}
	}
}
