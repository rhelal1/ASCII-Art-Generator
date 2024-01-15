package ascii

import (
	"fmt"
	"strings"
)

// function that handles the align operation flag
func PrintWithJustify(Words [][]string, column, row, count int) string {
	//initialize variables textSize as an empty string, spaceSize as 1, and width as the width of the console minus 1.
	textSize, spaceSize, width := "", 1, width()-1

	// iterates over each word in the Words slice and appends the value of Words[column][3] (which represents the word size) to textSize.
	for column := 0; column < len(Words); column++ {
		textSize += Words[column][3]
	}

	//checks if the length of textSize is less than width. If it is, the difference between width and len(textSize) is assigned to
	if len(textSize) < width {
		spaceSize = width - len(textSize)
	}

	//If align is "right" and column is 0, return a formatted string with spaceSize number of spaces followed by Words[column][row]
	if align == "right" && column == 0 {
		return fmt.Sprintf(strings.Repeat(" ", spaceSize) + Words[column][row])

		//If align is "center" and column is 0, return a formatted string with spaceSize/2 number of spaces followed by Words[column][row].
	} else if align == "center" && column == 0 {
		return fmt.Sprintf(strings.Repeat(" ", spaceSize/2) + Words[column][row])

		//If align is "justify", column is not 0, and specific indices of Words[column] contain specific values (indicating they need justification)
	} else if align == "justify" && column != 0 && Words[column][3] == "      " && Words[column][5] == "      " && Words[column][7] == "      " {
		if count == 0 { //if count is 0, it re-assign it to 1
			count = 1
		}

		//returns a formatted string with spaceSize/count number of spaces followed by Words[column][row].
		return fmt.Sprintf(strings.Repeat(" ", spaceSize/count) + Words[column][row])
	}

	//f none of the above conditions match (or align == "left"), return Words[column][row] without any formatting.
	return fmt.Sprintf(Words[column][row])
}
