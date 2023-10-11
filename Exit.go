package hangman

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func Exit(clear string) {
	fmt.Print(clear)
	content, _ := ioutil.ReadFile("affichage/goodbye.txt")
	fmt.Println(string(content))
	time.Sleep(3 * time.Second)
	fmt.Print(clear)
	os.Exit(0)
}
