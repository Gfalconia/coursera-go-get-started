// Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

// Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

// Submit your source code for the program,
// “findian.go”.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Printf("Please enter string to find: ")
	fmt.Scan(&str)
	str = strings.ToLower(str)
	if str[0:1] == "i" && str[len(str)-1:] == "n" && strings.Contains(str, "a") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
