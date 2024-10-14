package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var floatPoint float64
	// Print user instruction
	fmt.Printf("Hello, please enter a float: ")
	// Input statement
	fmt.Scan(&floatPoint)
	// Convert the float to a string
	float_string := strconv.FormatFloat(floatPoint, 'f', -1, 64)
	// split on the "." and return the first slice index to the screen
	fmt.Println((strings.Split(float_string, "."))[0])

}
