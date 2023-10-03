package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func loadWords() []string {
	arrayOfWords := []string{}
	inter := ""
	data, err := ioutil.ReadFile("words.txt")
	if err != nil {
		panic(err)
	}
	for _, char := range data {
		if char == '\n' {
			arrayOfWords = append(arrayOfWords, inter)
			inter = ""
			continue
		}
		inter = inter + string(char)
	}
	return arrayOfWords
}

func randomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(words))
	return words[random]
}

func displayWord(word string) string {
	display := ""
	for i := 0; i < len(word); i++ {
		display = display + "_"
	}
	for i := 0; i < (len(word))/2-1; i++ {
		randomNumber := rand.Intn(len(word) - 1)
		for string(display[randomNumber]) != "_" {
			randomNumber = rand.Intn(len(word) - 1)
		}
		display = display[:randomNumber] + string(word[randomNumber]) + display[randomNumber+1:]
	}
	return display
}
func askUser(display string, word string, life int, indexHangman int, failed_letter string) (string, int, int, string) {
	var input string
	fmt.Print("Choose : ")
	fmt.Scanln(&input)
	fmt.Print("\033[H\033[2J")
	display, life, indexHangman, failed_letter = isPresent(strings.ToUpper(input), word, display, life, indexHangman, failed_letter)
	return display, life, indexHangman, failed_letter
}

func isPresent(letter string, word string, display string, life int, indexHangman int, failed_letter string) (string, int, int, string) {
	isFind := false
	for i, char := range word {
		if i == len(word)-1 && string(char) == letter {
			display = display[:i] + letter
			isFind = true
			continue
		}
		if string(char) == letter {
			display = display[:i] + letter + display[i+1:]
			isFind = true
		}
	}
	if !isFind {
		if contains(failed_letter, letter) {
			fmt.Println("Not present in the word", life, "attemps remaining")
			fmt.Println("You already tried this letter")
			fmt.Println("False letters you have already tried: ", failed_letter)
			if life <= 10 {
				indexHangman = displayHangman(life, indexHangman)
			}
			return display, life, indexHangman, failed_letter
		}
		if !contains(failed_letter, letter) {
			if failed_letter == "" {
				failed_letter = failed_letter + letter
			} else {
				failed_letter = failed_letter + "," + letter
			}
		}
		life -= 1
		fmt.Println("Not present in the word", life, "attemps remaining")
		fmt.Println("False letters you have already tried: ", failed_letter)
		if life < 10 {
			indexHangman += 7
			indexHangman = displayHangman(life, indexHangman)
		}
		return display, life, indexHangman, failed_letter
	} else {
		fmt.Println("Present in the word", life, "attemps remaining")
		fmt.Println("False letters you have already tried: ", failed_letter)
		if life <= 10 {
			indexHangman = displayHangman(life, indexHangman)
		}
		return display, life, indexHangman, failed_letter
	}
}

func wordFind(word string, display string) bool {
	if !contains(display, "_") {
		fmt.Println(display)
		fmt.Println("Congrats ")
		return true
	}
	return false
}

func contains(inter string, x string) bool {
	for i := 0; i < len(inter); i++ {
		if x == string(inter[i]) {
			return true
		}
	}
	return false
}

func displayHangman(life int, indexHangman int) int {
	file, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	if life == 10 {
		for i := 0; i < 7; i++ {
			fmt.Println(lines[i])
		}
	} else {
		for i := indexHangman; i < indexHangman+7; i++ {
			fmt.Println(lines[i])
		}
	}
	return indexHangman
}

func main() {
	life := 10
	failed_letter := ""
	indexHangman := 0
	words := loadWords()
	word := randomWord(words)
	display := displayWord(word)
	fmt.Println("Good Luck, you have 10 attemps.")
	for !wordFind(word, display) && life > 0 {
		fmt.Println(display)
		display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter)
	}
	if life == 0 {
		fmt.Println("You lose, the good words was : ", word)
	}
}
