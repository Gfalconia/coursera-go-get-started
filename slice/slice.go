// Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

// Submit your source code for the program,
// “slice.go”.

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	var slc []int
	for true {
		fmt.Printf("Enter integer to add it to slice or X to end: ")
		fmt.Scan(&input)
		if input != "X" {
			sti, _ := strconv.Atoi(input)
			slc = append(slc, sti)
			sort.Ints(slc)
			fmt.Println(slc)
		} else {
			break
		}

	}

}
