package hangman

import (
	"fmt"
	"io/ioutil"
	"time"
)

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
