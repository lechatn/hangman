package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func menu() {
	reader := bufio.NewReader(os.Stdin)
	spaces := strings.Repeat(" ", 50)
	for {
		fmt.Println(spaces, "                       \033[35m||\033[0m                       ")
		fmt.Println(spaces, "                       \033[35m||\033[0m                       ")
		fmt.Println(spaces, "\033[35m----------------------Menu----------------------\033[0m")
		fmt.Println(spaces, "\033[35m|\033[0m    [1] : Play to hangman in "+"\033[36mfr\033[0m"+"en"+"\033[31mch\033[0m"+"           \033[35m|\033[0m")
		fmt.Println(spaces, "\033[35m|\033[0m    [2] : Play to hangman in "+"\033[31men\033[0m"+"gli"+"\033[34msh\033[0m"+"          \033[35m|\033[0m")
		fmt.Println(spaces, "\033[35m|\033[0m    [3] : Play to hangman in "+"\033[32mit\033[0m"+"ali"+"\033[31man\033[0m"+"          \033[35m|\033[0m")
		fmt.Println(spaces, "\033[35m|\033[0m    [4] : Play to hangman in "+"\033[31msp\033[0m"+"\033[33mani\033[0m"+"\033[31msh\033[0m"+"          \033[35m|\033[0m")
		fmt.Println(spaces, "\033[35m|\033[0m    [5] : Play to hangman with french citys   \033[35m|\033[0m")
		fmt.Println(spaces, "\033[35m|\033[0m    [6] : Play to hangman with countrys       \033[35m|\033[0m")
		fmt.Println(spaces, "\033[35m|\033[0m    [7] : Play to hangman with capitals       \033[35m|\033[0m")
		fmt.Println(spaces, "\033[35m|\033[0m    [8] : Play to hangman with sports         \033[35m|\033[0m")
		fmt.Println(spaces, "\033[35m|\033[0m    [9] : Leave the game                      \033[35m|\033[0m")
		fmt.Println(spaces, "\033[35m------------------------------------------------\033[0m")
		fmt.Print(spaces, " Choose an option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		switch option {
		case "1":
			fmt.Print("\033[H\033[2J")
			life := 10
			failed_letter := ""
			indexHangman := 0
			words := loadWords("words.txt")
			word := randomWord(words)
			display := displayWord(word)
			fmt.Println(spaces, "Game mode : \033[36mfr\033[0m"+"en"+"\033[31mch\033[0m words")
			fmt.Println(spaces, "Good Luck, you have 10 attemps.")
			for !wordFind(word, display) && life > 0 {
				fmt.Println(spaces, display)
				display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter)
			}
			if life == 0 {
				fmt.Print("\033[H\033[2J")
				fmt.Println(spaces, "You lose, the good words was : ", word)
			}
		case "2":
			fmt.Print("\033[H\033[2J")
			life := 10
			failed_letter := ""
			indexHangman := 0
			words := loadWords("english.txt")
			word := randomWord(words)
			display := displayWord(word)
			fmt.Println(spaces, "Game mode : \033[31men\033[0m"+"gli"+"\033[34msh\033[0m words")
			fmt.Println(spaces, "Good Luck, you have 10 attemps.")
			for !wordFind(word, display) && life > 0 {
				fmt.Println(spaces, display)
				display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter)
			}
			if life == 0 {
				fmt.Print("\033[H\033[2J")
				fmt.Println(spaces, "You lose, the good words was : ", word)
			}
		case "3":
			fmt.Print("\033[H\033[2J")
			life := 10
			failed_letter := ""
			indexHangman := 0
			words := loadWords("italiano.txt")
			word := randomWord(words)
			display := displayWord(word)
			fmt.Println(spaces, "Game mode : \033[32mit\033[0m"+"ali"+"\033[31man\033[0m words")
			fmt.Println(spaces, "Good Luck, you have 10 attemps.")
			for !wordFind(word, display) && life > 0 {
				fmt.Println(spaces, display)
				display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter)
			}
			if life == 0 {
				fmt.Print("\033[H\033[2J")
				fmt.Println(spaces, "You lose, the good words was : ", word)
			}
		case "4":
			fmt.Print("\033[H\033[2J")
			life := 10
			failed_letter := ""
			indexHangman := 0
			words := loadWords("espanol.txt")
			word := randomWord(words)
			display := displayWord(word)
			fmt.Println(spaces, "Game mode : \033[31msp\033[0m"+"\033[33mani\033[0m"+"\033[31msh\033[0m"+" words")
			fmt.Println(spaces, "Good Luck, you have 10 attemps.")
			for !wordFind(word, display) && life > 0 {
				fmt.Println(spaces, display)
				display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter)
			}
			if life == 0 {
				fmt.Print("\033[H\033[2J")
				fmt.Println(spaces, "You lose, the good words was : ", word)
			}
		case "5":
			fmt.Print("\033[H\033[2J")
			life := 10
			failed_letter := ""
			indexHangman := 0
			words := loadWords("villes_france.txt")
			word := randomWord(words)
			display := displayWord(word[:len(word)-1])
			fmt.Println(spaces, "Game mode : French citys")
			fmt.Println(spaces, "Good Luck, you have 10 attemps.")
			for !wordFind(word, display) && life > 0 {
				fmt.Println(spaces, display)
				display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter)
			}
			if life == 0 {
				fmt.Print("\033[H\033[2J")
				fmt.Println(spaces, "You lose, the good words was : ", word)
			}
		case "6":
			fmt.Print("\033[H\033[2J")
			life := 10
			failed_letter := ""
			indexHangman := 0
			words := loadWords("pays.txt")
			word := randomWord(words)
			display := displayWord(word[:len(word)-1])
			fmt.Println(spaces, "Game mode : Countrys")
			fmt.Println(spaces, "Good Luck, you have 10 attemps.")
			for !wordFind(word, display) && life > 0 {
				fmt.Println(spaces, display)
				display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter)
			}
			if life == 0 {
				fmt.Print("\033[H\033[2J")
				fmt.Println(spaces, "You lose, the good words was : ", word)
			}
		case "7":
			fmt.Print("\033[H\033[2J")
			life := 10
			failed_letter := ""
			indexHangman := 0
			words := loadWords("capital.txt")
			word := randomWord(words)
			display := displayWord(word[:len(word)-1])
			fmt.Println(spaces, "Game mode : Capitals")
			fmt.Println(spaces, "Good Luck, you have 10 attemps.")
			for !wordFind(word, display) && life > 0 {
				fmt.Println(spaces, display)
				display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter)
			}
			if life == 0 {
				fmt.Print("\033[H\033[2J")
				fmt.Println(spaces, "You lose, the good words was : ", word)
			}
		case "8":
			fmt.Print("\033[H\033[2J")
			life := 10
			failed_letter := ""
			indexHangman := 0
			words := loadWords("sports.txt")
			word := randomWord(words)
			display := displayWord(word)
			fmt.Println(spaces, "Game mode : Sports ")
			fmt.Println(spaces, "Good Luck, you have 10 attemps.")
			for !wordFind(word, display) && life > 0 {
				fmt.Println(spaces, display)
				display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter)
			}
			if life == 0 {
				fmt.Print("\033[H\033[2J")
				fmt.Println(spaces, "You lose, the good words was : ", word)
			}
		case "9":
			fmt.Print("\033[H\033[2J")
			content, _ := ioutil.ReadFile("goodbye.txt")
			fmt.Println(string(content))
			time.Sleep(3 * time.Second)
			fmt.Print("\033[H\033[2J")
			os.Exit(0)
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println(spaces, "Invalid option")
		}
	}
}

func loadWords(fichier string) []string {
	arrayOfWords := []string{}
	inter := ""
	data, err := ioutil.ReadFile(fichier)
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
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(word); i++ {
		if !contains(alphabet, string(word[i])) {
			display = display + string(word[i])
			continue
		}
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
	spaces := strings.Repeat(" ", 50)
	fmt.Print(spaces, " Choose : ")
	fmt.Scanln(&input)
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstvuwxyz"
	for len(input) > 1 || !contains(alphabet, input) {
		fmt.Println(spaces, "Invalid character")
		fmt.Print(spaces, " Choose : ")
		fmt.Scanln(&input)
	}
	fmt.Print("\033[H\033[2J")
	display, life, indexHangman, failed_letter = isPresent(strings.ToUpper(input), word, display, life, indexHangman, failed_letter)
	return display, life, indexHangman, failed_letter
}

func isPresent(letter string, word string, display string, life int, indexHangman int, failed_letter string) (string, int, int, string) {
	spaces := strings.Repeat(" ", 50)
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
			fmt.Println(spaces, "Not present in the word", life, "attemps remaining")
			fmt.Println(spaces, "You already tried this letter")
			fmt.Println(spaces, "False letters you have already tried: ", failed_letter)
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
		fmt.Println(spaces, "Not present in the word", life, "attemps remaining")
		fmt.Println(spaces, "False letters you have already tried: ", failed_letter)
		if life < 10 {
			indexHangman += 7
			indexHangman = displayHangman(life, indexHangman)
		}
		return display, life, indexHangman, failed_letter
	} else {
		fmt.Println(spaces, "Present in the word", life, "attemps remaining")
		fmt.Println(spaces, "Right letters you have already tried: ", failed_letter)
		if life <= 10 {
			indexHangman = displayHangman(life, indexHangman)
		}
		return display, life, indexHangman, failed_letter
	}
}

func wordFind(word string, display string) bool {
	spaces := strings.Repeat(" ", 50)
	if !contains(display, "_") {
		fmt.Println(spaces, display)
		fmt.Print("\033[H\033[2J")
		content, _ := ioutil.ReadFile("congrats.txt")
		fmt.Println(spaces, string(content))
		fmt.Println(spaces, "You find the word: ", word)
		time.Sleep(5 * time.Second)
		fmt.Print("\033[H\033[2J")
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
	spaces := strings.Repeat(" ", 50)
	file, err := ioutil.ReadFile("hangman.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
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

func main() {
	fmt.Print("\033[H\033[2J")
	menu()
}
