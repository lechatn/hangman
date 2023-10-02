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
	for _,char := range data {
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
		randomNumber := rand.Intn(len(word)-1)
		for string(display[randomNumber]) != "_" {
			randomNumber = rand.Intn(len(word)-1)
		}
		display = display[:randomNumber] + string(word[randomNumber]) + display[randomNumber+1:]
	}
	return display
}
func askUser(display string, word string,life int) (string,int) {
	var input string
	fmt.Println("Enter a letter: ")
	fmt.Scanln(&input)
	display,life = isPresent(strings.ToUpper(input),word,display,life)
	return display,life
}

func isPresent(letter string, word string,display string,life int) (string,int){
	isFind := false
	for i, char := range word {
		if string(char) == letter {
			display = display[:i] + letter + display[i+1:]
			isFind = true
		}
	}
	if !isFind {
		fmt.Println("Wrong letter")
		life -= 1
		return display,life
	} else {
		fmt.Println(display)
		return display,life
	}
}

func main() {
	life := 10
	words := loadWords()
	word := randomWord(words)
	display := displayWord(word)
	fmt.Println(display)
	for {
		display,life = askUser(display,word,life)
		fmt.Println(life)
	}
	
}