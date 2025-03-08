package main

import "fmt"

func main() {
	var length float64 = 5.0 // user input
	var width float64 = 3.0   // user input

	area := length * width
	perimeter := 2 * (length + width)

	fmt.Println("The area of the rectangle is", area)
	fmt.Println("The perimeter of the rectangle is", perimeter)
}
