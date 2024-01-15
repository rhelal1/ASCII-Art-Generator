package ascii

import (
	"fmt"
	"os"
	"strings"
)

var (
	colorFlag       = false      //flag to indicate if the --color= option is provided
	outputFlag      = false      //flag to indicate if the --output= option is provided
	reverseFlag     = false      //flag to indicate if the --reverse= option is provided
	fsFlag          = false      //flag to indicate if the font option is provided
	alignFlag       = false      //flag to indicate if the --align option is provided
	stringFlag      = false      //flag to indicate if the input text is provided
	colorWLflag     = false      //flag to indicate if the color option is applied to specific letter
	colorW2Lflag    = false      //flag to indicate if the color option is applied to specific 2 letters
	userText        = ""         //The user-provided text
	reverseFileName = ""         //The name of the file that contains the text for reverse operation
	fontFileName    = "standard" //The font's name, by default, it standard
	align           = "left"     //The alignment option for the output, by default, it left
	letterToColored = ""         //The letter or word to be colored.
	color           = "\033[0m"  //The color code
	fileOutputName  = ""         //The name of the output file
	validation      = "color"    //The type of operation being performed, depending on flags usage
)

// function to validate the user input as command-line arguments and determine the type of
// operation that needs to be performed with the right parameters given by the user
func Validation() (string, []string) {

	for i := 1; i < len(os.Args); i++ { // Iterate over the command-line arguments

		// Check if the colorW2L operation is being performed
		if (len(os.Args) == 6 || len(os.Args) == 7) && strings.Index(os.Args[1], "--color=") == 0 && strings.Index(os.Args[3], "--color=") == 0 && i == 1 {
			validation, colorW2Lflag = "colorW2L", true            //set the colorW2Lflag as true, and validation as "colorW2L"
			LetterExistnce(os.Args[5], os.Args[4])                 //check the letter/word existence in the text that is provided by the user
			i += 3                                                 //skip 3 arguments, to check to last 2 only
			color, letterToColored = ColorValdation(1), os.Args[2] //validate the color given by the user

			// Check if the output file option is provided
		} else if strings.Index(os.Args[i], "--output=") == 0 && !stringFlag && !outputFlag {
			//set the outputFlag as true, validation as "output", and filename after removeing the leading characters ("--output=")
			validation, outputFlag, fileOutputName = "output", true, os.Args[i][9:]
			os.Remove(fileOutputName) //remove the file, if it exists

			//validate the file name (the extension should be .txt) and filename should not be empty
			if (strings.Index(os.Args[i], ".txt") == -1) || len(os.Args[i]) != strings.Index(os.Args[i], ".txt")+4 {
				fmt.Println("The file name is incorrect")
				Error()
			}

			// Check if the color option is provided
		} else if strings.Index(os.Args[i], "--color=") == 0 && !stringFlag && !colorFlag && !colorW2Lflag {
			colorFlag, color = true, ColorValdation(i) //validate the color, and set the color flag as true

			// Check if the alignment option is provided
		} else if strings.Index(os.Args[i], "--align=") == 0 && !stringFlag && !alignFlag {
			//set the alignFlag as true, and align after removeing the leading characters ("--align=")
			align, alignFlag = strings.ToLower(os.Args[i][8:]), true

			//if align not equal to justify, left, right or center, call Error() function
			if !(align == "justify" || align == "left" || align == "right" || align == "center") {
				fmt.Println("align is incorrect, you can use (left, right, center, justify) only")
				Error()
			}

			// Check if the reverse option is provided
		} else if strings.Index(os.Args[i], "--reverse=") == 0 && !stringFlag && !reverseFlag {
			//set the reverseFlag as true, validation as Reverse, and reverseFileName after removeing the leading characters ("--reverse=")
			validation, reverseFlag, reverseFileName = "Reverse", true, os.Args[i][10:]

			//validate the file name (the extension should be .txt)
			if (strings.Index(os.Args[i], ".txt") == -1) || len(os.Args[i]) != strings.Index(os.Args[i], ".txt")+4 {
				fmt.Println("The file name is incorrect")
				Error()
			} else if len(os.Args) != 2 { //if arguments are not equal to 2 in reverse operation, call Error() function
				Error()
			}

			// Check if the color option is applied to specific letter
		} else if colorFlag && i+1 < len(os.Args) && !colorWLflag && os.Args[i+1] != "standard" && os.Args[i+1] != "shadow" && os.Args[i+1] != "thinkertoy" {
			//set the colorWLflag as true, and letterToColored as the argument
			colorWLflag, letterToColored = true, os.Args[i]

		} else if !stringFlag && !reverseFlag {
			//set stringFlag as true, and remove any tab '\\t' from the userText
			userText, stringFlag = strings.ReplaceAll(os.Args[i], "\\t", "   "), true
			//validate the userText should contain only printable characters (between 32-126 ASCII-code)
			for char := 0; char < len(userText); char++ {
				if userText[char] > 126 || userText[char] < 32 {
					fmt.Println("ERROR: ascii letters only")
					os.Exit(0)
				}
			}

			//check if the font option is provided (should be after the user provides the Text
		} else if stringFlag && !fsFlag {
			//set fsFlag flag as true, and fontFileName as the argument
			fontFileName, fsFlag = strings.ToLower(os.Args[i]), true

			//if fontFileName is not equal to standard, shadow or thinkertoy, call Error() function
			if fontFileName != "standard" && fontFileName != "shadow" && fontFileName != "thinkertoy" {
				Error()
			}

		} else { //if no condition is met, call Error() function
			Error()

		}
	}

	if alignFlag {
		userText = strings.Join(strings.Fields(userText), " ") // Ensure proper alignment of the userText, by removing extra spaces
	}

	// Check if color is applied to specific word/letter (letterToColored) is exists in the text (userText)
	if colorW2Lflag || colorWLflag {
		LetterExistnce(userText, letterToColored)
	}

	// Check if the input text or reverse operation is provided, otherwise call Error() function
	if !stringFlag && !reverseFlag {
		Error()
	}

	if outputFlag && colorFlag { //if color and output flags are being used together, print a warning message
		fmt.Println("The colored text can't be print inside the file")
	}

	// Split the user text into individual lines
	WordsInArr := strings.Split(userText, "\\n")

	// Remove the last element if it only contains a new line character
	if OnlyContainsNewLine(userText) {
		WordsInArr = WordsInArr[:len(WordsInArr)-1]
	}

	return validation, WordsInArr
}

// function to validate the color given by the user using ColorValdation function
func ColorValdation(index int) string {
	color := CheckColor(os.Args[index]) //set color as the argument in the given index
	if color == "NOTVALID" {            //if color = "NOTVALID", print an error message and call Error() function
		fmt.Println("color does not exist")
		Error()
	}
	return color
}

// function to checks if the string b exists within the string a
// If b is not found in a or if b is an empty string, it prints an error message and calls the Error function.
func LetterExistnce(a, b string) {
	if strings.Index(a, b) == -1 || len(b) == 0 {
		fmt.Println("letters to be colored does not exist in the word")
		Error()
	}
}

// function to prints an error message indicating the usage of the program and an example of the command. It exits the program.
func Error() {
	fmt.Println("Usage: go run main.go [OPTION] [STRING] [BANNER]\nEX: go run main.go --align=right --output=fileName.txt --color=red \"something\" shadow")
	os.Exit(0)
}

// function to check if the string (text) contains only new lines
func OnlyContainsNewLine(text string) bool {
	for i := 0; i < len(text); i++ { //iterates over each character in the string text
		//If the current character and the next character form the escape sequence "\n" (new line) and it is
		//not the last two characters in the string, it skips the next character by incrementing i
		if string(text[i]) == "\\" && string(text[i+1]) == "n" && i != len(text)-3 {
			i++
		} else { //If any character is encountered that does not match the escape sequence "\n", it returns false
			return false
		}
	}

	//If the loop completes without encountering any non-matching character, it returns true,
	//indicating that the string only contains new line characters
	return true
}
