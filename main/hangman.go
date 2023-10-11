package main

import ( // Import of the packages
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"hangman"
)

func menu() { // Display of command menu
	reader := bufio.NewReader(os.Stdin)
	var find bool
	spaces := strings.Repeat(" ", 50)
	clear := "\033[H\033[2J"
	red := "\033[31m" // We define all the color we want to use
	blue := "\033[34m"
	black := "\033[30m"
	magenta := "\033[35m"
	cyan := "\033[36m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset_color := "\033[0m"
	score := 0
	win_series := 1
	for { //Define the principal menu
		fmt.Println(spaces, "                        "+magenta+"||"+reset_color+"                       ")
		fmt.Println(spaces, "                        "+magenta+"||"+reset_color+"                       ")
		fmt.Println(spaces, magenta, "----------------------Menu----------------------", reset_color, spaces, "score : ", score)
		fmt.Println(spaces, magenta, "|"+reset_color+"    [1] : Language                            "+magenta+"|", reset_color)
		fmt.Println(spaces, magenta, "|"+reset_color+"    [2] : Others                              "+magenta+"|", reset_color)
		fmt.Println(spaces, magenta, "|"+reset_color+"    [3] : Leave the game                      "+magenta+"|", reset_color)
		fmt.Println(spaces, magenta, "------------------------------------------------", reset_color)
		fmt.Println(spaces, "\033[3m Created by Guillaume, Arthur and Noé", reset_color)
		fmt.Println()
		fmt.Println()
		fmt.Print(spaces, "  Choose an option: ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		fmt.Print(clear)
		switch option {
		case "1":
			finish := false
			for !finish { // Define the submenu named "Language"
				fmt.Println(spaces, "                         "+magenta+"||"+reset_color+"                       ")
				fmt.Println(spaces, "                         "+magenta+"||"+reset_color+"                       ")
				fmt.Println(spaces, magenta+"----------------------Language----------------------"+reset_color, spaces, "score : ", score)
				fmt.Println(spaces, magenta+"|"+reset_color+"    [1] : Play to hangman in "+cyan+"fr"+reset_color+"en"+red+"ch"+reset_color+"               "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"    [2] : Play to hangman in "+red+"en"+reset_color+"gli"+blue+"sh"+reset_color+"              "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"    [3] : Play to hangman in "+green+"it"+reset_color+"ali"+red+"an"+reset_color+"              "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"    [4] : Play to hangman in "+red+"sp"+reset_color+yellow+"ani"+reset_color+red+"sh"+reset_color+"              "+magenta+"|", reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"    [5] : Play to hangman in "+red+"por"+reset_color+green+"tugu"+reset_color+red+"ese"+reset_color+"           "+magenta+"|", reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"    [6] : Play to hangman in "+black+"ge"+reset_color+red+"rm"+reset_color+yellow+"an"+reset_color+"               "+magenta+"|", reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"    [7] : Back to menu                            "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"    [8] : Leave the game                          "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"----------------------------------------------------"+reset_color)
				fmt.Print(spaces, " Choose an option: ")

				option, _ := reader.ReadString('\n')
				option = strings.TrimSpace(option)
				fmt.Print(clear)
				life := 10
				failed_letter := ""
				indexHangman := 0
				switch option {
				case "1":
					words := loadWords("base_de_donnée/words.txt") // Reading of all the words in the database
					word := randomWord(words)                      // Word choice for playing
					display := displayWord(word)                   // Display of "_" and starting letters
					game_mode := "Game mode : \033[36mfr" + reset_color + "en" + red + "ch" + reset_color + " words" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "2":
					words := loadWords("base_de_donnée/english.txt") // Reading of all the words in the database
					word := randomWord(words)                        // Word choice for playing
					display := displayWord(word)                     // Display of "_" and starting letters
					game_mode := "Game mode : \033[31men" + reset_color + "gli" + blue + "sh" + reset_color + " words" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "3":
					words := loadWords("base_de_donnée/italiano.txt") // Reading of all the words in the database
					word := randomWord(words)                         // Word choice for playing
					display := displayWord(word)                      // Display of "_" and starting letters
					game_mode := "Game mode : \033[32mit" + reset_color + "ali" + red + "an" + reset_color + " words" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "4":
					words := loadWords("base_de_donnée/espanol.txt") // Reading of all the words in the database
					word := randomWord(words)                        // Word choice for playing
					display := displayWord(word)                     // Display of "_" and starting letters
					game_mode := "Game mode : \033[31msp" + reset_color + yellow + "ani" + reset_color + red + "sh" + reset_color + " words" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "5":
					words := loadWords("base_de_donnée/portugais.txt") // Reading of all the words in the database
					word := randomWord(words)                          // Word choice for playing
					display := displayWord(word)                       // Display of "_" and starting letters
					game_mode := "Game mode : \033[31mpor" + reset_color + green + "tugu" + reset_color + red + "ese" + reset_color + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "6":
					words := loadWords("base_de_donnée/allemand.txt") // Reading of all the words in the database
					word := randomWord(words)                         // Word choice for playing
					display := displayWord(word)                      // Display of "_" and starting letters
					game_mode := "Game mode : \033[30mge" + reset_color + red + "rm" + reset_color + yellow + "an" + reset_color + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "8":
					hangman.Exit(clear)
				case "7":
					fmt.Print(clear)
					finish = true
				default:
					fmt.Print(clear)
					fmt.Println(spaces, "Invalid option")
				}
			}

		case "2":
			finish := false
			for !finish { // Define the submenu "Others"
				fmt.Println(spaces, "                         "+magenta+"||"+reset_color+"                        ")
				fmt.Println(spaces, "                         "+magenta+"||"+reset_color+"                        ")
				fmt.Println(spaces, magenta+"-----------------------Others-----------------------"+reset_color, spaces, "score : ", score)
				fmt.Println(spaces, magenta+"|"+reset_color+"     [1] : Play to hangman with french citys      "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"     [2] : Play to hangman with countrys          "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"     [3] : Play to hangman with capitals          "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"     [4] : Play to hangman with sports            "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"     [5] : Play to hangman with brands            "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"     [6] : Play to hangman with foods             "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"     [7] : Play to hangman with drinks            "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"     [8] : Play to hangman with league of legends "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"     [9] : Back to menu                           "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"|"+reset_color+"     [10] : Leave the game                        "+magenta+"|"+reset_color)
				fmt.Println(spaces, magenta+"----------------------------------------------------"+reset_color)
				fmt.Print(spaces, " Choose an option: ")

				option, _ := reader.ReadString('\n')
				option = strings.TrimSpace(option)
				fmt.Print(clear)
				life := 10
				failed_letter := ""
				indexHangman := 0
				switch option {
				case "1":
					words := loadWords("base_de_donnée/villes_france.txt") // Reading of all the words in the database
					word := randomWord(words)                              // Word choice for playing
					display := displayWord(word[:len(word)-1])
					game_mode := "Game mode : French citys" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "2":
					words := loadWords("base_de_donnée/pays.txt") // Reading of all the words in the database
					word := randomWord(words)                     // Word choice for playing
					display := displayWord(word[:len(word)-1])
					game_mode := "Game mode : Countrys" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "3":
					words := loadWords("base_de_donnée/capital.txt") // Reading of all the words in the database
					word := randomWord(words)                        // Word choice for playing
					display := displayWord(word[:len(word)-1])
					game_mode := "Game mode : Capitals" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "4":
					words := loadWords("base_de_donnée/sports.txt") // Reading of all the words in the database
					word := randomWord(words)                       // Word choice for playing
					display := displayWord(word)                    // Display of "_" and starting letters
					game_mode := "Game mode : Sports " + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "5":
					words := loadWords("base_de_donnée/marque.txt") // Reading of all the words in the database
					word := randomWord(words)                       // Word choice for playing
					display := displayWord(word)                    // Display of "_" and starting letters
					game_mode := "Game mode : Brands " + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "6":
					words := loadWords("base_de_donnée/food.txt") // Reading of all the words in the database
					word := randomWord(words)                     // Word choice for playing
					display := displayWord(word)                  // Display of "_" and starting letters
					game_mode := "Game mode : Food " + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "7":
					words := loadWords("base_de_donnée/boissons.txt") // Reading of all the words in the database
					word := randomWord(words)                         // Word choice for playing
					display := displayWord(word)                      // Display of "_" and starting letters
					game_mode := "Game mode : Drinks " + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "8":
					words := loadWords("base_de_donnée/lol.txt") // Reading of all the words in the database
					word := randomWord(words)                    // Word choice for playing
					display := displayWord(word)                 // Display of "_" and starting letters
					game_mode := "Game mode : League of legends " + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = hangman.WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = askUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						win_series = hangman.Lose(win_series, clear, spaces, word)
					}
				case "10":
					hangman.Exit(clear)
				case "9":
					fmt.Print(clear)
					finish = true
				default:
					fmt.Print(clear)
					fmt.Println(spaces, "Invalid option")
				}
			}
		case "3":
			hangman.Exit(clear)
		default:
			fmt.Print(clear)
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

func randomWord(words []string) string { // Function who choose a random word in the array of words
	rand.Seed(time.Now().UnixNano()) // We initialize a new seed based on the actual hour
	random := rand.Intn(len(words))
	return words[random]
}

func displayWord(word string) string { // Function who create the game, we display some "_" and certaine letters of the word choosen
	display := ""
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letter_display := 0
	for i := 0; i < len(word); i++ {
		if !hangman.Contains(alphabet, string(word[i])) {
			display = display + string(word[i])
			letter_display++
			continue
		}
		display = display + "_"
	}
	fmt.Print(letter_display)
	for i := 0; i < (len(word))/2-1-letter_display; i++ { // At the beggining, we display len(word)/2-1 characters of the good word
		randomNumber := rand.Intn(len(word) - 1)
		for string(display[randomNumber]) != "_" {
			randomNumber = rand.Intn(len(word) - 1)
		}
		display = display[:randomNumber] + string(word[randomNumber]) + display[randomNumber+1:]
	}
	return display
}
func askUser(display string, word string, life int, indexHangman int, failed_letter string, game_mode string) (string, int, int, string) { // Function who verify if there is an error when the user give him a letter
	clear := "\033[H\033[2J"
	var input string
	spaces := strings.Repeat(" ", 50)
	fmt.Print(spaces, " Choose : ")
	fmt.Scanln(&input)
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstvuwxyz"
	for len(input) > 1 || !hangman.Contains(alphabet, input) { //	We display an error message if the user give us an invalid character
		fmt.Println(spaces, "Invalid character")
		fmt.Print(spaces, " Choose : ")
		fmt.Scanln(&input)
	}
	fmt.Print(clear)
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
		if hangman.Contains(failed_letter, letter) {
			fmt.Println(spaces, game_mode)
			fmt.Println(spaces, "Not present in the word", life, "attemps remaining")
			fmt.Println(spaces, "You already tried this letter")
			fmt.Println(spaces, "False letters you have already tried: ", failed_letter) //	We remind to the user the false letters that he already tried
			if life <= 10 {
				indexHangman = hangman.DisplayHangman(life, indexHangman)
			}
			return display, life, indexHangman, failed_letter
		}
		if !hangman.Contains(failed_letter, letter) {
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
			indexHangman = hangman.DisplayHangman(life, indexHangman)
		}
		return display, life, indexHangman, failed_letter
	} else { // if the letter is in the word, we said this to the user and we display the present part of José
		fmt.Println(spaces, game_mode)
		fmt.Println(spaces, "Present in the word", life, "attemps remaining")
		fmt.Println(spaces, "Right letters you have already tried: ", failed_letter)
		if life <= 10 {
			indexHangman = hangman.DisplayHangman(life, indexHangman)
		}
		return display, life, indexHangman, failed_letter
	}
}

func main() { // We define the main function who launch the game
	clear := "\033[H\033[2J"
	fmt.Print(clear)
	menu()
}
