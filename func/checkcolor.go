package ascii

import (
	"fmt"
	"strconv"
	"strings"
)

// function to validate the color given by the user
func CheckColor(userValue string) string {
	var RGB [3]int
	var err error

	// Convert user input to lowercase and remove the leading characters ("--color=")
	userValue = strings.ToLower(userValue[8:])

	// Map of color names to ANSI escape sequences
	ANSIColors := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"purple":  "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"black":   "\033[30m",
		"orange":  "\033[38;5;208m",
		"\033[0m": "\033[0m"}

	// Validate the color
	for color, value := range ANSIColors {
		// Check if the user input matches a predefined color in ANSI Map
		if color == userValue {
			return value
		}
	}

	// Check if the user input represents an RGB color in the format "rgb(x, y, z)"
	if strings.Index(userValue, "rgb") == 0 && userValue[len(userValue)-1] == ')' {
		userValue := strings.ReplaceAll(userValue, " ", "")

		// Extract the RGB values from the input string
		RGBValues := (strings.Split(userValue[4:len(userValue)-1], ","))

		// Validate the RGB values
		for i := 0; i < len(RGBValues); i++ {
			RGB[i], err = strconv.Atoi(RGBValues[i])
			if err != nil || len(RGBValues) != 3 || RGB[i] < 0 || RGB[i] > 255 {
				return "NOTVALID"
			}
		}

		// Convert RGB values to ANSI escape sequence
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", RGB[0], RGB[1], RGB[2])

		// Check if the user input represents a color in hexadecimal format "#rrggbb"
	} else if strings.Index(userValue, "#") == 0 && len(userValue) == 7 {

		// Valid hexadecimal digit
		for char := 1; char <= 6; char++ {
			if !((userValue[char] >= '0' && userValue[char] <= '9') || (userValue[char] >= 'a' && userValue[char] <= 'f')) {
				return "NOTVALID"
			}
		}

		// Convert hexadecimal color value to RGB values, then ANSI escape sequence
		rgb, _ := strconv.ParseUint(userValue[1:], 16, 32)
		userValue = fmt.Sprintf("\033[38;2;%d;%d;%dm", int(rgb>>16&0xFF), int(rgb>>8&0xFF), int(rgb&0xFF))
		return userValue
	}

	//If the user input does not match any valid color format, return "NOTVALID"
	return "NOTVALID"
}
