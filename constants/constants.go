package constants

import (
	"bufio"
	"fmt"
	"os"
)

const UserFile = "../../data/users.json"
const PostFile = "../../data/posts.json"

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const Gray = "\033[37m"
const White = "\033[97m"

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
