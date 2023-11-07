package hangman

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

func DisplayHangman(life int, indexHangman int) int { // Function who display a part of José
	spaces := strings.Repeat(" ", 50)
	file, err := ioutil.ReadFile("affichage/hangman.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n") // We cut the file hangman.txt line by line
	if life == 10 {
		for i := 0; i < 7; i++ {
			fmt.Println(spaces, lines[i])
		}
	} else {
		for i := indexHangman; i < indexHangman+7; i++ {
			fmt.Println(spaces, lines[i])
		}
	}
	return indexHangman
}

func DisplayWord(word string) string { // Function who create the game, we display some "_" and certaine letters of the word choosen
	display := ""
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letter_display := 0
	for i := 0; i < len(word); i++ {
		if !Contains(alphabet, string(word[i])) {
			display = display + string(word[i])
			letter_display++
			continue
		}
		display = display + "_"
	}
	for i := 0; i < (len(word))/2-1-letter_display; i++ { // At the beggining, we display len(word)/2-1 characters of the good word
		randomNumber := rand.Intn(len(word) - 1)
		for string(display[randomNumber]) != "_" {
			randomNumber = rand.Intn(len(word) - 1)
		}
		display = display[:randomNumber] + string(word[randomNumber]) + display[randomNumber+1:]
	}
	return display
}
