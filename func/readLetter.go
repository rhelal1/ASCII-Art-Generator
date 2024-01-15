package ascii

import (
	"bufio"
	"fmt"
	"os"
)

// function converts the given letter in bytes as parameters into his ASCII representation
func ReadLetter(LetterInByte byte) []string {
	stop := 1                                    // Variable to control the stop condition for reading lines
	var LetterInASCII []string                   // Slice to store the lines of the letter
	letterLength := (int(LetterInByte)-32)*9 + 2 //Calculate the length of the letter based on the ASCII value of Text1

	ReadFile, err := os.Open("./fonts/" + fontFileName + ".txt") //Open the file with the specified fileName
	if err != nil {
		fmt.Println(err) //If the file can't be opened print the error and call the Error() function
		Error()
	}

	FileScanner := bufio.NewScanner(ReadFile) // Create a Scanner to read the file line by line

	for i := 1; FileScanner.Scan() && stop <= 8; i++ { // Iterate over each line of the file
		if i >= letterLength {
			stop++
			LetterInASCII = append(LetterInASCII, FileScanner.Text()) // Append the line to the Letter slice
		}
	}

	ReadFile.Close()     // Close the file
	return LetterInASCII //Return the slice containing the lines of the letter
}
