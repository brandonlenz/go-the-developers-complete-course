package main

import "fmt"

func main() {
	t := Triangle{
		width:  5,
		height: 3,
	}

	fmt.Println("Triangle area is:", t.getArea())

	s := Square{
		sideLength: 5,
	}

	fmt.Println("Square area is:", s.getArea())
}
