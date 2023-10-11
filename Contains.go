package hangman

func Contains(inter string, x string) bool { // Function who test if a string is present on another string
	for i := 0; i < len(inter); i++ {
		if x == string(inter[i]) {
			return true
		}
	}
	return false
}
