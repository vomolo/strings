package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the number of arguments is exactly 1
	if len(os.Args) != 2 {
		return
	}

	// Get the input string from the command line argument
	input := os.Args[1]

	// Slice to hold the words
	var words []string
	word := ""

	// Iterate over each character in the input string
	for i := 0; i < len(input); i++ {
		char := input[i]

		// Check if the character is a space
		if char == ' ' {
			// If we have accumulated a word, add it to the words slice
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			// Accumulate characters to form a word
			word += string(char)
		}
	}

	// Add the last word if it exists
	if word != "" {
		words = append(words, word)
	}

	// Check if there are no words
	if len(words) == 0 {
		return
	}

	// Create the output with three spaces between words
	output := ""
	for i, w := range words {
		output += w
		if i < len(words)-1 {
			output += "   " // Add three spaces between words
		}
	}

	// Print the output followed by a newline
	fmt.Println(output)
}
