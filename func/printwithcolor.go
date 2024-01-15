package ascii

import (
	"fmt"
	"os"
	"strings"
)

// function that handles the color operation flags
func PrintWithColor(Words [][]string, Text string, count int) {
	//call the CheckColors function with the Text argument to get the positions where colors need to be applied.
	positions := CheckColors(Text)

	//iterates over 8 times (from 0 to 7) for each line (row) and each width.
	for row := 0; row < 8; row++ {
		for column := 0; column < len(Words); column++ {
			//prints the color at the corresponding position from positions map, followed by the result of the PrintWithJustify function.
			fmt.Print(positions[column], PrintWithJustify(Words, column, row, count))
		}

		if row+1 != 8 { //prints a new line, except for the last iteration
			fmt.Println()
		}
	}

	fmt.Print("\033[0m", "\n") //prints an empty line and resetting the color. For formating only
}

// function that takes Text and returns a map of the text where the keys are integers representing positions and the values are strings representing colors.
func CheckColors(Text string) map[int]string {
	indexWcolor := make(map[int]string) //empty map to store the positions and colors

	for i := 0; i < len(Text); i++ { // iterates over each character in Text
		//if substring starting from the current position and with the length of letterToColored
		//matches letterToColored and letterToColored is not an empty string, enters the first if block.
		if i+len(letterToColored) <= len(Text) && Text[i:i+len(letterToColored)] == letterToColored && letterToColored != "" {

			//loops over the matched characters and sets the corresponding positions in indexWcolor map to the specified color
			for j := 0; j < len(letterToColored); j++ {
				indexWcolor[i] = color
			}

			//increments i by the length of letterToColored minus 1 to skip the matched characters in the next iteration
			i += len(letterToColored) - 1

			//if thevalidation string contains "2" (colorW2L) and the substring starting from the current
			// position and with the length of os.Args[4] (Letter 2) matches os.Args[4], enters the second if block.
		} else if strings.Contains(validation, "2") && i+len(letterToColored) <= len(Text) && Text[i:i+len(os.Args[4])] == os.Args[4] {

			//loops over the matched characters and sets the corresponding positions in indexWcolor
			// map to the color returned by the CheckColor function with os.Args[3] as an argument.
			for j := 0; j < len(os.Args[4]); j++ {
				indexWcolor[i] = CheckColor(os.Args[3])
			}

			//increments i by the length of os.Args[4] minus 1 to skip the matched characters in the next iteration
			i += len(os.Args[4]) - 1

			//If letterToColored is an empty string, it sets the current position in indexWcolor map to the specified color.
		} else if letterToColored == "" {
			indexWcolor[i] = color

			//If none of the above conditions match, it sets the current position in indexWcolor map to the
			// escape sequence "\033[0m" for resetting the color.
		} else {
			indexWcolor[i] = "\033[0m"
		}
	}

	return indexWcolor //returns the indexWcolor map.
}
