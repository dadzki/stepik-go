package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var runeStr []rune

	fmt.Scan(&str)

	for _, char := range str {
		if strings.Count(str, string(char)) == 1 {
			runeStr = append(runeStr, char)
		}
	}

	fmt.Println(string(runeStr))
}
