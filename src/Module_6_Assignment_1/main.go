/*

Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

*/

package main

import (
	"fmt"
)

func user_input() string {
	// Function used to gather user input then return it
	var temp_name string
	fmt.Scan(&temp_name)
	return temp_name

}

func main() {
	// Used to collect user details and return a map address
	fmt.Println("Please enter your first name:")
	first_name := user_input()
	fmt.Println("Please enter your last name:")
	last_name := user_input()
	fmt.Println("Please enter your address:")
	address := user_input()
	fmt.Println(first_name, last_name, address)

	full_name := first_name + " " + last_name

	// Create the map that will be used to contain all of the user details
	var userMap map[string]string
	userMap = make(map[string]string)
	// Begin adding values to the map
	userMap["Name"] = full_name

	userMap["address"] = address

	fmt.Println(userMap)

}
