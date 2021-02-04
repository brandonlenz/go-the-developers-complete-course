package main

import "fmt"

func main() {
	numbers := []int{}

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)

		if i%2 != 0 {
			fmt.Printf("%d is odd\n", i)
		} else {
			fmt.Printf("%d is even\n", i)
		}
	}
}
