package hangman

import (
	"fmt"
	"strings"
)

func AskUser(display string, word string, life int, indexHangman int, failed_letter string, game_mode string) (string, int, int, string) { // Function who verify if there is an error when the user give him a letter
	clear := "\033[H\033[2J"
	var input string
	spaces := strings.Repeat(" ", 50)
	fmt.Print(spaces, " Choose : ")
	fmt.Scanln(&input)
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstvuwxyz"
	for len(input) > 1 || !Contains(alphabet, input) { //	We display an error message if the user give us an invalid character
		fmt.Println(spaces, "Invalid character")
		fmt.Print(spaces, " Choose : ")
		fmt.Scanln(&input)
	}
	fmt.Print(clear)
	display, life, indexHangman, failed_letter = IsPresent(strings.ToUpper(input), word, display, life, indexHangman, failed_letter, game_mode)
	return display, life, indexHangman, failed_letter
}
