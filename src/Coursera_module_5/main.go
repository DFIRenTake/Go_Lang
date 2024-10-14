/*

Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	// When declaring a slice, unless you want to exlude the length\capacity, you have to add those items
	// To do that you need to use the make command.
	slice := make([]int, 0, 3)

	// Create the loop that will run until the user enter the letter X
	for {
		// Initialize the integer variable as a string
		var integer string
		// Print message requesting user input
		fmt.Printf("Please insert some info: ")
		// Use the scan function to store the user input into the integer memory space
		fmt.Scan(&integer)

		// Evaluate the user input to determine if the letter "x" has been provided, if not, continue the loop, if so break
		if integer == "X" {
			// Break the currently infinite loop
			break
		}

		// This action require that you include the new variable, and an index for the value you are inputting.
		i, _ := strconv.Atoi(integer)
		// Append the slice with the new integer
		slice = append(slice, i)

		// sort the integers
		sort.Ints(slice)
		fmt.Println(slice)

	}
}
