package main

import ( // Import of the packages
	"fmt"
	"hangman"
)

func main() { // We define the main function who launch the game
	clear := "\033[H\033[2J"
	fmt.Print(clear)
	hangman.Menu()
}
