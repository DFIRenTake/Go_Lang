package main

/*

Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!”
otherwise. The program should not be case-sensitive,
so it does not matter if the characters are upper-case or lower-case.


*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Please insert some info: ")
	// Take the input and set it equal to in
	in := bufio.NewReader(os.Stdin)
	// Set the input of the bufio equal to the string read
	line, _ := in.ReadString('\n')

	// Take the string, convert to lower, and remove space
	new_line := strings.ToLower(line)

	// use strings.index to find the position of the string indicated
	output1 := strings.Index(new_line, "i")
	output2 := strings.Index(new_line, "n")
	output3 := strings.Index(new_line, "a")
	// Determine the length of the string and -1
	true_len := len(strings.TrimSpace(new_line)) - 1

	// Use switch and case to determine if the criteria is met
	switch {
	case output1 == 0 && output2 == true_len && output3 != -1:
		fmt.Print("found")

	default:
		fmt.Print("not found")

	}

}
