package hangman

import (
	"io/ioutil"
	"math/rand"
	"time"
)

func LoadWords(fichier string) []string { // Define a fonction who create a aeeay of words with all the words of the database
	arrayOfWords := []string{}
	inter := ""
	data, err := ioutil.ReadFile(fichier)
	if err != nil {
		panic(err)
	}
	for _, char := range data { // We append all the words of the database in a variable
		if char == '\n' {
			arrayOfWords = append(arrayOfWords, inter)
			inter = ""
			continue
		}
		inter = inter + string(char)
	}
	return arrayOfWords
}

func RandomWord(words []string) string { // Function who choose a random word in the array of words
	rand.Seed(time.Now().UnixNano()) // We initialize a new seed based on the actual hour
	random := rand.Intn(len(words))
	return words[random]
}
