package main

import (
    "fmt"
    "io/ioutil"
)

func loadWords() []string {
	arrayOfWords := []string{}
	inter := ""
	data, err := ioutil.ReadFile("words.txt")
	if err != nil {
		panic(err)
	}
	for _,char := range data {
		inter = inter + string(char)
		if char == '\n' {
			arrayOfWords = append(arrayOfWords, inter)
			inter = ""
		}
	}
	arrayOfWords = append(arrayOfWords, inter)
	return arrayOfWords
}

func main() {
	words := loadWords()
	fmt.Println(words)
}