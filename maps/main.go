package main

import "fmt"

func printColors(c map[string]string) {
	for color, hexCode := range c {
		fmt.Println("The color", color, "has a hex code of", hexCode)
	}
}

func main() {

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "0000FF",
		"white": "FFFFFF",
	}

	printColors(colors)
}
