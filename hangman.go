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
	for i := 0; i < len(word)-1; i++ {
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
func askUser(display string, word string, life int, indexHangman int) (string, int , int) {
	var input string
	fmt.Print("Choose : ")
	fmt.Scanln(&input)
	display, life , indexHangman = isPresent(strings.ToUpper(input), word, display, life, indexHangman)
	return display, life , indexHangman
}

func isPresent(letter string, word string, display string, life int, indexHangman int) (string, int , int) {
	isFind := false
	for i, char := range word {
		if i == len(word)-1 && string(char) == letter{
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
		life -= 1
		fmt.Println("Not present in the word", life, "attemps remaining", "\n")
		if life < 10 {
			indexHangman +=7
			indexHangman = displayHangman(life, indexHangman)
		}
		return display, life , indexHangman
	} else {
		if life <10 {
			indexHangman = displayHangman(life, indexHangman)	
		}
		return display, life , indexHangman
	}
}

func wordFind(word string, display string) bool {
	if !contains(display, "_") {
		fmt.Println(display)
		fmt.Println("Congrats ", "\n")
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

func displayHangman(life int, indexHangman int) int{
	file, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for i := indexHangman; i < indexHangman+7; i++ {
		fmt.Println(lines[i])
	}
	return indexHangman
}

func main() {
	life := 10
	indexHangman := -7
	words := loadWords()
	word := randomWord(words)
	display := displayWord(word)
	fmt.Println("Good Luck, you have 10 attemps.")
	for !wordFind(word, display) && life > 0 {
		fmt.Println(display)
		display, life , indexHangman = askUser(display, word, life, indexHangman)
	}
	if life == 0 {
		fmt.Println("You lose, the good words was : ", word)
	}

}
