package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Extract text from standard.txt in a byteSlice type
	byteSlice, _ := os.ReadFile("standard.txt")

	// Convert the the byteslice into string type
	text := string(byteSlice)

	// Split the string into each ascii element
	slice := strings.Split(text, "\n\n")

	// Make sure the only one argument is to be entered
	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Println("Invalid input")
		return
	}

	// Create a variable from the argument input
	in := os.Args[1]

	for _, c := range in {
		if (c < 32 || c > 126 ) && c != 10 {
			fmt.Println("Invalid input")
			return
		}
	}

	// Handle the single new line case in the argument
	if in == "\\n" {
		fmt.Println()
		return
	}

	// handle the single empty string case in the argument
	if in == "" {
		return
	}

	// Split the input argument variable in with new lines if there are any
	inSplit := finalSplit(in)

	// Start iterating over a slice of strings separated by newline
	for _, c := range inSplit {
		// Create j and k counters for eliminating extra newlines
		j := 0
		k := 0
		// Start an inner loop for printing line by line
		for i := 0; i <= 7; i++ {
			// Start iterating over each character of a string slice element
			for _, e := range c {
				a := strings.Split(string(slice[e-32]), "\n")
				// Handle the case of first line of space character
				if i == 0 && e == ' ' {
					fmt.Print((a[1]))
				} else {
					fmt.Print((a[i]))
				}
			}
			// Handling the case of extra new lines
			if c != "" {
				fmt.Println()
			}
			k++

		}
		j++
		// Handling the case of extra new lines
		if c == "" || (j == len(inSplit)-1 && k == 7) {
			fmt.Println()
		}

	}
}

//Handles cases of visible and unvisible newlines
func finalSplit(str string) []string {
	word := ""
	var final []string
	slice := strings.Split(str, "\\n")
	for _, el := range slice {
		for _, char := range el {
			if char == 10 {
				final = append(final, word)
				word = ""
				continue
			} else {
				word += string(char)
			}
		}
		if word == "" {
			final = append(final, word)
		} else {
			final = append(final, word)
			word = ""
		}
	}
	return final
}
