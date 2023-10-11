package hangman

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func DisplayHangman(life int, indexHangman int) int { // Function who display a part of José
	spaces := strings.Repeat(" ", 50)
	file, err := ioutil.ReadFile("affichage/hangman.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n") // We cut the file hangman.txt line by line
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
