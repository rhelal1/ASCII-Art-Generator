package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var AllLeters [][]string //2D array to store the ASCII letters

// function that handles the reverse operation flag
func PrintWithReverse() {

	var fileLine [][]string //2D array to store the lines read from the file
	var words []string      //string array to store the converted words

	ReadFile, err := os.Open(reverseFileName) // Open the file for reading

	if err != nil {
		fmt.Println(err) // Print the error if there's an issue opening the file
		Error()
	}

	FileScanner := bufio.NewScanner(ReadFile) // Create a scanner to read the file line by line

	for FileScanner.Scan() { // Read each line from the file and split it into individual characters
		fileLine = append(fileLine, strings.Split(FileScanner.Text(), ""))
	}

	ReadFile.Close() // Close the file after reading

	if len(fileLine)%8 != 0 { // Check if the number of lines is divisible by 8, otherwise display an error
		Error()
	}

	AppendLetters() // Load the ASCII letters from the font files

	// Process the lines in blocks of 8
	for i := 0; i < len(fileLine); i += 8 {
		var Line [][]string // Store the current block of 8 lines

		for j := i; j < i+8 && j < len(fileLine); j++ { // Extract the block of 8 lines
			Line = append(Line, fileLine[j])
		}

		for len(Line[0]) > 0 {
			n, find := 0, false

			// Iterate over the ASCII letters and check if the current characters match any letter
			for i := 0; i < len(AllLeters); i++ {
				if i == 94 || i == 188 {
					n++
				}
				if CheckTheLetter(Line, AllLeters[i]) {
					// Convert the matched letter to a string and append to words
					words = append(words, string(i+32-(94*n)))
					// Trim the processed characters from the lines
					Terimming(len(AllLeters[i][0]), len(Line[0]), Line)
					find = true
					break
				}
			}
			if !find { // If no match is found, display an error
				Error()
			}
		}

		// Append a newline character after processing each block of lines
		words = append(words, "\n")
	}

	// Print the converted words
	for i := 0; i < len(words)-1; i++ {
		fmt.Print(words[i])
	}

	// Print an empty line after printing the words
	fmt.Println()
}

// Check if the characters in Line match the characters in AllLeters
func CheckTheLetter(Line [][]string, AllLeters []string) bool {
	for n := 0; n < len(AllLeters[0]); n++ {
		for row := 0; row < len(Line); row++ {
			// If any character doesn't match, return false
			if len(Line[row]) == 0 || len(AllLeters[0]) > len(Line[0]) || Line[row][n] != string(AllLeters[row][n]) {
				return false
			}
		}
	}

	return true //All characters match, return true
}

// function to trim the processed characters from the lines
func Terimming(startIndex, endIndex int, Line [][]string) {

	for i := 0; i < len(Line); i++ {
		if len(Line[i]) >= endIndex {
			Line[i] = Line[i][startIndex:endIndex] // Trim the line to remove the processed characters

		} else { //If the line is shorter than endIndex, display an error
			Error()
		}
	}
}

// function to load the ASCII letters from the font files and append them to AllLeters
func AppendLetters() {
	files := [3]string{"standard", "shadow", "thinkertoy"}
	for j := 0; j < len(files); j++ { //loop over the fonts
		for i := 32; i < 126; i++ {
			fontFileName = files[j]
			// Read and append each ASCII letter to AllLeters
			AllLeters = append(AllLeters, ReadLetter(byte(i)))
		}
	}
}
