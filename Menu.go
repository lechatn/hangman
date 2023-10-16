package hangman

import ( // Import of the packages
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Menu() { // Display of command menu
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
		fmt.Println(green, "██╗"+reset_color+"░░░░░"+green+"███████╗████████╗██╗"+reset_color+"░"+green+"██████╗            ", reset_color, magenta, "|"+reset_color+"    [1] : Language                            "+magenta+"|", green, "              ██████╗"+reset_color+"░"+green+"██╗"+reset_color+"░░░░░░"+green+"█████╗"+reset_color+"░"+green+"██╗"+reset_color+"░░░"+green+"██╗")
		fmt.Println(green, "██║"+reset_color+"░░░░░"+green+"██╔════╝╚══██╔══╝╚█║██╔════╝            ", reset_color, magenta, "|"+reset_color+"    [2] : Others                              "+magenta+"|", green, "              ██╔══██╗██║"+reset_color+"░░░░░"+green+"██╔══██╗╚██╗"+reset_color+"░"+green+"██╔╝")
		fmt.Println(green, "██║"+reset_color+"░░░░░"+green+"█████╗"+reset_color+"░░░░░"+green+"██║"+reset_color+"░░░░"+green+"╚╝╚█████╗"+reset_color+"░            ", reset_color, magenta, "|"+reset_color+"    [3] : Leave the game                      "+magenta+"|", green, "              ██████╔╝██║"+reset_color+"░░░░░"+green+"███████║"+reset_color+"░"+green+"╚████╔╝"+reset_color+"░")
		fmt.Println(green, "██║"+reset_color+"░░░░░"+green+"██╔══╝"+reset_color+"░░░░░"+green+"██║"+reset_color+"░░░░░░░"+green+"╚═══██╗"+reset_color+"            ", reset_color, magenta, "------------------------------------------------", green, "              ██╔═══╝"+reset_color+"░"+green+"██║"+reset_color+"░░░░░"+green+"██╔══██║"+reset_color+"░░"+green+"╚██╔╝"+reset_color+"░░", reset_color)
		fmt.Println(green, "███████╗███████╗"+reset_color+"░░░"+green+"██║"+reset_color+"░░░░░░"+green+"██████╔╝"+reset_color+"            ", reset_color, "\033[3m Created by Guillaume, Arthur and Noé", green, "                          ██║"+reset_color+"░░░░░"+green+"███████╗██║"+reset_color+"░░"+green+"██║"+reset_color+"░░░"+green+"██║"+reset_color+"░░░", reset_color)
		fmt.Println(green, "╚══════╝╚══════╝"+reset_color+"░░░"+green+"╚═╝"+reset_color+"░░░░░░"+green+"╚═════╝"+reset_color+"░                                                                             ", green, "╚═╝"+reset_color+"░░░░░"+green+"╚══════╝╚═╝"+reset_color+"░░"+green+"╚═╝"+reset_color+"░░░"+green+"╚═╝"+reset_color+"░░░", reset_color)
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
				fmt.Println("         █░░ ▄▀█ █▄░█ █▀▀ █░█ ▄▀█ █▀▀ █▀▀         ", magenta+"|"+reset_color+"    [4] : Play to hangman in "+red+"sp"+reset_color+yellow+"ani"+reset_color+red+"sh"+reset_color+"              "+magenta+"|", reset_color+"         █░░ ▄▀█ █▄░█ █▀▀ █░█ ▄▀█ █▀▀ █▀▀")
				fmt.Println("         █▄▄ █▀█ █░▀█ █▄█ █▄█ █▀█ █▄█ ██▄         ", magenta+"|"+reset_color+"    [5] : Play to hangman in "+red+"por"+reset_color+green+"tugu"+reset_color+red+"ese"+reset_color+"           "+magenta+"|", reset_color+"         █▄▄ █▀█ █░▀█ █▄█ █▄█ █▀█ █▄█ ██▄")
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
					words := LoadWords("base_de_donnée/words.txt") // Reading of all the words in the database
					word := RandomWord(words)                      // Word choice for playing
					display := DisplayWord(word)                   // Display of "_" and starting letters
					game_mode := "Game mode : \033[36mfr" + reset_color + "en" + red + "ch" + reset_color + " words" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "2":
					words := LoadWords("base_de_donnée/english.txt") // Reading of all the words in the database
					word := RandomWord(words)                        // Word choice for playing
					display := DisplayWord(word)                     // Display of "_" and starting letters
					game_mode := "Game mode : \033[31men" + reset_color + "gli" + blue + "sh" + reset_color + " words" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "3":
					words := LoadWords("base_de_donnée/italiano.txt") // Reading of all the words in the database
					word := RandomWord(words)                         // Word choice for playing
					display := DisplayWord(word)                      // Display of "_" and starting letters
					game_mode := "Game mode : \033[32mit" + reset_color + "ali" + red + "an" + reset_color + " words" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "4":
					words := LoadWords("base_de_donnée/espanol.txt") // Reading of all the words in the database
					word := RandomWord(words)                        // Word choice for playing
					display := DisplayWord(word)                     // Display of "_" and starting letters
					game_mode := "Game mode : \033[31msp" + reset_color + yellow + "ani" + reset_color + red + "sh" + reset_color + " words" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "5":
					words := LoadWords("base_de_donnée/portugais.txt") // Reading of all the words in the database
					word := RandomWord(words)                          // Word choice for playing
					display := DisplayWord(word)                       // Display of "_" and starting letters
					game_mode := "Game mode : \033[31mpor" + reset_color + green + "tugu" + reset_color + red + "ese" + reset_color + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "6":
					words := LoadWords("base_de_donnée/allemand.txt") // Reading of all the words in the database
					word := RandomWord(words)                         // Word choice for playing
					display := DisplayWord(word)                      // Display of "_" and starting letters
					game_mode := "Game mode : \033[30mge" + reset_color + red + "rm" + reset_color + yellow + "an" + reset_color + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "8":
					Exit(clear)
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
				fmt.Println("               █▀█ ▀█▀ █░█ █▀▀ █▀█ █▀             ", magenta+"|"+reset_color+"     [5] : Play to hangman with brands            "+magenta+"|"+reset_color, "              █▀█ ▀█▀ █░█ █▀▀ █▀█ █▀")
				fmt.Println("               █▄█ ░█░ █▀█ ██▄ █▀▄ ▄█             ", magenta+"|"+reset_color+"     [6] : Play to hangman with foods             "+magenta+"|"+reset_color, "              █▄█ ░█░ █▀█ ██▄ █▀▄ ▄█")
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
					words := LoadWords("base_de_donnée/villes_france.txt") // Reading of all the words in the database
					word := RandomWord(words)                              // Word choice for playing
					display := DisplayWord(word[:len(word)-1])
					game_mode := "Game mode : French citys" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "2":
					words := LoadWords("base_de_donnée/pays.txt") // Reading of all the words in the database
					word := RandomWord(words)                     // Word choice for playing
					display := DisplayWord(word[:len(word)-1])
					game_mode := "Game mode : Countrys" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "3":
					words := LoadWords("base_de_donnée/capital.txt") // Reading of all the words in the database
					word := RandomWord(words)                        // Word choice for playing
					display := DisplayWord(word[:len(word)-1])
					game_mode := "Game mode : Capitals" + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "4":
					words := LoadWords("base_de_donnée/sports.txt") // Reading of all the words in the database
					word := RandomWord(words)                       // Word choice for playing
					display := DisplayWord(word)                    // Display of "_" and starting letters
					game_mode := "Game mode : Sports " + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "5":
					words := LoadWords("base_de_donnée/marque.txt") // Reading of all the words in the database
					word := RandomWord(words)                       // Word choice for playing
					display := DisplayWord(word)                    // Display of "_" and starting letters
					game_mode := "Game mode : Brands " + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "6":
					words := LoadWords("base_de_donnée/food.txt") // Reading of all the words in the database
					word := RandomWord(words)                     // Word choice for playing
					display := DisplayWord(word)                  // Display of "_" and starting letters
					game_mode := "Game mode : Food " + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "7":
					words := LoadWords("base_de_donnée/boissons.txt") // Reading of all the words in the database
					word := RandomWord(words)                         // Word choice for playing
					display := DisplayWord(word)                      // Display of "_" and starting letters
					game_mode := "Game mode : Drinks " + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "8":
					words := LoadWords("base_de_donnée/lol.txt") // Reading of all the words in the database
					word := RandomWord(words)                    // Word choice for playing
					display := DisplayWord(word)                 // Display of "_" and starting letters
					game_mode := "Game mode : League of legends " + spaces + "score : " + strconv.Itoa(score)
					fmt.Println(spaces, game_mode)
					fmt.Println(spaces, "Good Luck, you have 10 attemps.")
					for !find && life > 0 {
						find, score = WordFind(word, display, score, win_series, life)
						if find {
							win_series++ // If the word is found, the number of winning series increases, and the score rises more quickly
							find = false
							break
						}
						fmt.Println(spaces, display)
						display, life, indexHangman, failed_letter = AskUser(display, word, life, indexHangman, failed_letter, game_mode)
					}
					if life == 0 {
						time.Sleep(1 * time.Second)
						win_series = Lose(win_series, clear, spaces, word)
					}
				case "10":
					Exit(clear)
				case "9":
					fmt.Print(clear)
					finish = true
				default:
					fmt.Print(clear)
					fmt.Println(spaces, "Invalid option")
				}
			}
		case "3":
			Exit(clear)
		default:
			fmt.Print(clear)
			fmt.Println(spaces, "Invalid option")

		}
	}
}
