package main

import ( // Import of the packages
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func menu() { // Display of command menu
	reader := bufio.NewReader(os.Stdin)
	spaces := strings.Repeat(" ", 50)
	red := "\033[31m"
	blue := "\033[34m"
	black := "\033[30m"
	magenta := "\033[35m"
	cyan := "\033[36m"
	green := "\033[32m"
	reset_color := "\033[0m"
	for { //Define the principal menu
		fmt.Println(spaces, "                       \033[35m||\033[0m                       ")
		fmt.Println(spaces, "                       \033[35m||\033[0m                       ")
		fmt.Println(spaces, magenta, "----------------------Menu----------------------", reset_color)
		fmt.Println(spaces, magenta, "|\033[0m    [1] : Language                            \033[35m|", reset_color)
		fmt.Println(spaces, magenta, "|\033[0m    [2] : Others                              \033[35m|", reset_color)
		fmt.Println(spaces, magenta, "|\033[0m    [3] : Leave the game                      \033[35m|", reset_color)
		fmt.Println(spaces, magenta, "------------------------------------------------", reset_color)
		fmt.Println(spaces, "\033[3mCreated by Guillaume, Arthur and Noé", reset_color)
		fmt.Println()
		fmt.Println()
		fmt.Print(spaces, " Choose an option: ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		fmt.Print("\033[H\033[2J")
		switch option {
		case "1":
			finish := false
			for !finish { // Define the submenu named "Language"
				fmt.Println(spaces, "                         \033[35m||\033[0m                       ")
				fmt.Println(spaces, "                         \033[35m||\033[0m                       ")
				fmt.Println(spaces, magenta+"----------------------Language----------------------"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [1] : Play to hangman in "+cyan+"fr"+reset_color+"en"+red+"ch"+reset_color+"               \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [2] : Play to hangman in "+red+"en"+reset_color+"gli"+blue+"sh"+reset_color+"              \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [3] : Play to hangman in "+green+"it"+reset_color+"ali"+red+"an"+reset_color+"              \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [4] : Play to hangman in "+red+"sp"+reset_color+"\033[33mani"+reset_color+red+"sh"+reset_color+"              \033[35m|", reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [5] : Play to hangman in "+red+"por"+reset_color+green+"tugu"+reset_color+red+"ese"+reset_color+"           \033[35m|", reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [6] : Play to hangman in "+black+"ge"+reset_color+red+"rm"+reset_color+"\033[33man"+reset_color+"               \033[35m|", reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [7] : Back to menu                            \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [8] : Leave the game                          \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"----------------------------------------------------"+reset_color)
				fmt.Print(spaces, " Choose an option: ")

				option, _ := reader.ReadString('\n')
				option = strings.TrimSpace(option)
				fmt.Print("\033[H\033[2J")
				life := 10
				failed_letter := ""
				indexHangman := 0
				switch option {
				case "1":
					words := loadWords("base_de_donnée/words.txt")
					word := randomWord(words)
					display := displayWord(word)
					game_mode := "Game mode : \033[36mfr" + reset_color + "en" + red + "ch\033[0m words"
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "2":
					words := loadWords("base_de_donnée/english.txt")
					word := randomWord(words)
					display := displayWord(word)
					game_mode := "Game mode : \033[31men" + reset_color + "gli" + blue + "sh\033[0m words"
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "3":
					words := loadWords("base_de_donnée/italiano.txt")
					word := randomWord(words)
					display := displayWord(word)
					game_mode := "Game mode : \033[32mit" + reset_color + "ali" + red + "an\033[0m words"
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "4":
					words := loadWords("base_de_donnée/espanol.txt")
					word := randomWord(words)
					display := displayWord(word)
					game_mode := "Game mode : \033[31msp" + reset_color + "\033[33mani" + reset_color + red + "sh" + reset_color + " words"
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "5":
					words := loadWords("base_de_donnée/portugais.txt")
					word := randomWord(words)
					display := displayWord(word)
					game_mode := "Game mode : \033[31mpor" + reset_color + green + "tugu" + reset_color + red + "ese" + reset_color
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "6":
					words := loadWords("base_de_donnée/allemand.txt")
					word := randomWord(words)
					display := displayWord(word)
					game_mode := "Game mode : \033[30mge" + reset_color + red + "rm" + reset_color + "\033[33man" + reset_color
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "8":
					fmt.Print("\033[H\033[2J")
					content, _ := ioutil.ReadFile("affichage/goodbye.txt")
					fmt.Println(string(content))
					time.Sleep(3 * time.Second)
					fmt.Print("\033[H\033[2J")
					os.Exit(0)
				case "7":
					fmt.Print("\033[H\033[2J")
					finish = true
				default:
					fmt.Print("\033[H\033[2J")
					fmt.Println(spaces, "Invalid option")
				}
			}

		case "2":
			finish := false
			for !finish { // Define the submenu "Others"
				fmt.Println(spaces, "                        \033[35m||\033[0m                        ")
				fmt.Println(spaces, "                        \033[35m||\033[0m                        ")
				fmt.Println(spaces, magenta+"----------------------Others----------------------"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [1] : Play to hangman with french citys     \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [2] : Play to hangman with countrys         \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [3] : Play to hangman with capitals         \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [4] : Play to hangman with sports           \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [5] : Play to hangman with brands           \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [6] : Play to hangman with foods            \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [7] : Play to hangman with drinks           \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [8] : Back to menu                          \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"|\033[0m    [9] : Leave the game                        \033[35m|"+reset_color)
				fmt.Println(spaces, magenta+"--------------------------------------------------"+reset_color)
				fmt.Print(spaces, " Choose an option: ")

				option, _ := reader.ReadString('\n')
				option = strings.TrimSpace(option)
				fmt.Print("\033[H\033[2J")
				life := 10
				failed_letter := ""
				indexHangman := 0
				switch option {
				case "1":
					words := loadWords("base_de_donnée/villes_france.txt")
					word := randomWord(words)
					display := displayWord(word[:len(word)-1])
					game_mode := "Game mode : French citys"
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "2":
					words := loadWords("base_de_donnée/pays.txt")
					word := randomWord(words)
					display := displayWord(word[:len(word)-1])
					game_mode := "Game mode : Countrys"
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "3":
					words := loadWords("base_de_donnée/capital.txt")
					word := randomWord(words)
					display := displayWord(word[:len(word)-1])
					game_mode := "Game mode : Capitals"
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "4":
					words := loadWords("base_de_donnée/sports.txt")
					word := randomWord(words)
					display := displayWord(word)
					game_mode := "Game mode : Sports "
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "5":
					words := loadWords("base_de_donnée/marque.txt")
					word := randomWord(words)
					display := displayWord(word)
					game_mode := "Game mode : Brands "
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "6":
					words := loadWords("base_de_donnée/food.txt")
					word := randomWord(words)
					display := displayWord(word)
					game_mode := "Game mode : Food "
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "7":
					words := loadWords("base_de_donnée/boissons.txt")
					word := randomWord(words)
					display := displayWord(word)
					game_mode := "Game mode : Drinks "
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !wordFind(word, display) && life > 0 {
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						fmt.Print("\033[H\033[2J")
						fmt.Println(spaces, "You lose, the good words was : ", word)
					}
				case "9":
					fmt.Print("\033[H\033[2J")
					content, _ := ioutil.ReadFile("affichage/goodbye.txt")
					fmt.Println(string(content))
					time.Sleep(3 * time.Second)
					fmt.Print("\033[H\033[2J")
					os.Exit(0)
				case "8":
					fmt.Print("\033[H\033[2J")
					finish = true
				default:
					fmt.Print("\033[H\033[2J")
					fmt.Println(spaces, "Invalid option")
				}
			}
		case "3":
			fmt.Print("\033[H\033[2J")
			content, _ := ioutil.ReadFile("affichage/goodbye.txt")
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
func loadWords(fichier string) []string { // Define a fonction who create a aeeay of words with all the words of the database
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

func randomWord(words []string) string { // Function who choose a random word in the array of words
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(words))
	return words[random]
}

func displayWord(word string) string { // Function who create the game, we display some "_" and certaine letters of the word choosen
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
func askUser(display string, word string, life int, indexHangman int, failed_letter string, game_mode string) (string, int, int, string) { // Function who verify if there is an error when the user give him a letter
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
	display, life, indexHangman, failed_letter = isPresent(strings.ToUpper(input), word, display, life, indexHangman, failed_letter, game_mode)
	return display, life, indexHangman, failed_letter
}

func isPresent(letter string, word string, display string, life int, indexHangman int, failed_letter string, game_mode string) (string, int, int, string) { // Function who verify if the letter is present in the word
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
	if !isFind { // if the letter is not in the word, we said this to the user and display the next part of José
		if contains(failed_letter, letter) {
			fmt.Println(spaces, game_mode)
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
		fmt.Println(spaces, game_mode)
		fmt.Println(spaces, "Not present in the word", life, "attemps remaining")
		fmt.Println(spaces, "False letters you have already tried: ", failed_letter)
		if life < 10 {
			indexHangman += 7
			indexHangman = displayHangman(life, indexHangman)
		}
		return display, life, indexHangman, failed_letter
	} else { // if the letter is in the word, we said this to the user and we display the present part of José
		fmt.Println(spaces, game_mode)
		fmt.Println(spaces, "Present in the word", life, "attemps remaining")
		fmt.Println(spaces, "Right letters you have already tried: ", failed_letter)
		if life <= 10 {
			indexHangman = displayHangman(life, indexHangman)
		}
		return display, life, indexHangman, failed_letter
	}
}

func wordFind(word string, display string) bool { // Function who verify if the user have find the word and display the message of congrats
	spaces := strings.Repeat(" ", 50)
	if !contains(display, "_") {
		fmt.Println(spaces, display)
		fmt.Print("\033[H\033[2J")
		content, _ := ioutil.ReadFile("affichage/congrats.txt")
		fmt.Println(spaces, string(content))
		fmt.Println(spaces, "You find the word: ", word)
		time.Sleep(5 * time.Second)
		fmt.Print("\033[H\033[2J")
		return true
	}
	return false
}

func contains(inter string, x string) bool { // Function who test if a string is present on another string
	for i := 0; i < len(inter); i++ {
		if x == string(inter[i]) {
			return true
		}
	}
	return false
}

func displayHangman(life int, indexHangman int) int { // Function who display a part of José
	spaces := strings.Repeat(" ", 50)
	file, err := ioutil.ReadFile("affichage/hangman.txt")
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
