package hangman

import (
	"fmt"
	"io/ioutil"
	"os"
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

func Lose(win_series int, clear string, spaces string, word string) int {
	fmt.Print(clear)
	if win_series != 1 {
		win_series = 1 // If you lose, the win counter resets to 1
	}
	content, _ := ioutil.ReadFile("affichage/loose.txt")
	fmt.Println(spaces, string(content))
	fmt.Println(spaces, "The good words was : ", word)
	time.Sleep(3 * time.Second)
	fmt.Print(clear)
	return win_series
}

func Exit(clear string) {
	fmt.Print(clear)
	content, _ := ioutil.ReadFile("affichage/goodbye.txt")
	fmt.Println(string(content))
	time.Sleep(3 * time.Second)
	fmt.Print(clear)
	os.Exit(0)
}
