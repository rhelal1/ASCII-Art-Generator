package ascii

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// function to check the width of the terminal with the user text input after converting it to ASCII representation
func CheckTextSizeWithWidth(word []string) bool {
	width := width() // Get the width of the terminal window

	for _, char := range word { // Iterate over each word in the slice
		wordWfont := ""
		for _, char2 := range char { // Iterate over each character in the word
			wordWfont += ReadLetter(byte(char2))[0] // Read the letter corresponding to the character and append it to wordWfont
		}
		if len(wordWfont) > width { // Check if the length of the wordWfont applied exceeds the terminal width
			return false // Return false if it exceeds the width
		}
	}
	return true // Return true if all words fit within the terminal width
}

func width() int {
	cmd := exec.Command("stty", "size") // Create a command to retrieve the terminal size
	cmd.Stdin = os.Stdin                // Set the standard input of the command to the current process's standard input
	out, err := cmd.Output()            // Execute the command and capture the output
	if err != nil {
		log.Fatal(err) // If there's an error executing the command, log the error and exit the program
	}

	// Split the output by space and extract the second part, which represents the terminal width
	width, err := strconv.Atoi(strings.Split(strings.TrimSpace(string(out)), " ")[1])
	if err != nil {
		log.Fatal(err) // If there's an error converting the width to an integer, log the error and exit the program
	}
	return width // Return the terminal width as an integer
}
