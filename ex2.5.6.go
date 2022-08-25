package main

import (
	"fmt"
	"unicode"
)

func validate(p string) bool {
	runePass := []rune(p)
	if len(runePass) < 5 {
		return false
	}

	for _, v := range runePass {
		if !unicode.Is(unicode.Latin, v) && !unicode.Is(unicode.Digit, v) {
			return false
		}
	}

	return true
}

func main() {
	var pass string

	fmt.Scan(&pass)

	if validate(pass) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}

}
