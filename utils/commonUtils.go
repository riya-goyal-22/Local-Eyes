package utils

import (
	"bufio"
	"fmt"
	"os"
)

func PromptInput(prompt string) string {

	// Create a new Scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Display the prompt message
	fmt.Print(prompt)

	// Read the next line of input
	scanner.Scan()

	// Get the text from the scanner
	input := scanner.Text()

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return input
}
