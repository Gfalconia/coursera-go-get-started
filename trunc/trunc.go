// Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

// Submit your source code for the program, “trunc.go”.
package main

import (
	"fmt"
)

func main() {
	var number float32
	fmt.Print("Type a number: ")
	fmt.Scan(&number)
	fmt.Printf("Your truncated number is: %d", int(number))
}
