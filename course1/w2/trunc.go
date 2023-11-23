package main

import (
	"fmt"
)

func main() {
	var number float64

	fmt.Println("Enter a floating point number: ")
	fmt.Scan(&number)
	fmt.Println("Truncated integer: ", int(number))
}
