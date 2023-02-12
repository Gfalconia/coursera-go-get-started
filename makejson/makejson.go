// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

// Submit your source code for the program,
// “makejson.go”.

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	var person Person
	fmt.Println("Please enter your name: ")
	fmt.Scan(&person.Name)
	fmt.Println("Please enter your address: ")
	fmt.Scan(&person.Address)
	jsonOut, _ := json.Marshal(person)
	fmt.Println(string(jsonOut))
}
