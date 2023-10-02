package main

import (
    "fmt"
    "io/ioutil"
	"math/rand"
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
	for i := 0; i < len(word)/2-1; i++ {
		randomNumber := rand.Intn(len(word)-1)
		display = display[:randomNumber] + string(word[randomNumber]) + display[randomNumber+1:]
	}
	return display
}

func main() {
	words := loadWords()
	word := randomWord(words)
	display := displayWord(word)
	fmt.Println(display)
}