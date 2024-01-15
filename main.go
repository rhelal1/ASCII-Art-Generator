package main

import (
	"ascii/func"
	"fmt"
	"os"
	"strings"
)

func main() {
	//call Validation() function, to validate the user input and determine the required operation
	validation, WordsInArr := ascii.Validation()

	//Check if the validation contains "R" for reverse flag
	if strings.Contains(validation, "R") {
		ascii.PrintWithReverse()
		os.Exit(0)
	}

	for l := 0; l < len(WordsInArr); l++ { //Iterate over each word in WordsInArr
		var Words [][]string
		count := strings.Count(WordsInArr[l], " ") //count the number of spaces to use it in justify

		for j := 0; j < len(WordsInArr[l]); j++ { //Read each letter in the word and store it in Words as ASCII representation
			Words = append(Words, ascii.ReadLetter(WordsInArr[l][j]))
		}

		//Check if the validation contains "output" for writing to a file
		if strings.Contains(validation, "output") {
			ascii.WriteFile(Words, count)

			//Check if the validation contains "color" for printing with color. By default will use print with color
			//with terminal default color if not specified by the user
		} else if strings.Contains(validation, "color") {

			if !(ascii.CheckTextSizeWithWidth(WordsInArr)) { //Check if the text size in ASCII representation is within the width of the terminal
				fmt.Println("length of the ASCII Text is higher than the length of the terminal, try shorter text.")
				os.Exit(0)

			} else if len(WordsInArr[l]) != 0 { //Print the ASCII art with color if the word is not empty
				ascii.PrintWithColor(Words, WordsInArr[l], count)

			} else { // Print a new line if the word is empty
				fmt.Println()
			}
		}
	}
}
