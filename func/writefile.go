package ascii

import (
	"fmt"
	"os"
)

// function to print the ASCII representation of the text in a file given by the user with the '--output=' flag
func WriteFile(TheOutPutText [][]string, count int) {
	file, err := os.OpenFile(fileOutputName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err) // Print the error if there's an issue opening the file
		os.Exit(0)
	}

	for row := 0; row < 8 && len(TheOutPutText) != 0; row++ { // Iterate over each row (line) in the TheOutPutText
		for column := 0; column < len(TheOutPutText); column++ { // Iterate over each column (character) in the row
			file.WriteString(PrintWithJustify(TheOutPutText, column, row, count)) // Write the justified character to the file
		}

		if row+1 != 8 {
			file.WriteString("\n") // Write a newline character after each row, except the last one
		}
	}

	file.WriteString("\n") // Write an additional newline character after all rows are written

	file.Close() // Close the file
}
