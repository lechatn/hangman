package hangman

import (
)

func IsPresent(letter string, word string, display string, life int, indexHangman int, failed_letter string) (string, int, int, string) { // Function who verify if the letter is present in the word
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
		if Contains(failed_letter, letter) {
			if life <= 10 {
				indexHangman = DisplayHangman(life, indexHangman)
			}
			return display, life, indexHangman, failed_letter
		}
		if !Contains(failed_letter, letter) {
			if failed_letter == "" {
				failed_letter = failed_letter + letter
			} else {
				failed_letter = failed_letter + "," + letter
			}
		}
		life -= 1
		if life < 10 {
			indexHangman += 7
			indexHangman = DisplayHangman(life, indexHangman)
		}
		return display, life, indexHangman, failed_letter
	} else { // if the letter is in the word, we said this to the user and we display the present part of José
		if life <= 10 {
			indexHangman = DisplayHangman(life, indexHangman)
		}
		return display, life, indexHangman, failed_letter
	}
}

func Contains(inter string, x string) bool { // Function who test if a string is present on another string
	for i := 0; i < len(inter); i++ {
		if x == string(inter[i]) {
			return true
		}
	}
	return false
}
