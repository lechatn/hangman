package hangman

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func WordFind(word string, display string, score int, win_series int, life int) (bool, int) { // Function who verify if the user have find the word and display the message of congrats
	clear := "\033[H\033[2J"
	spaces := strings.Repeat(" ", 50)
	if !Contains(display, "_") {
		fmt.Println(spaces, display)
		fmt.Print(clear)
		life++
		score = score + life/2*win_series // We add to the score life/2*number of series
		content, _ := ioutil.ReadFile("affichage/congrats.txt")
		fmt.Println(spaces, string(content))
		fmt.Println(spaces, " You find the word: ", word)
		time.Sleep(5 * time.Second)
		fmt.Print(clear)
		return true, score
	}
	return false, score
}