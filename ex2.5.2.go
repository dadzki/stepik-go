package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)

	runeString := []rune(text)

	length := len(runeString)
	for i := 0; i < length/2; i++ {
		if runeString[i] != runeString[length-i-1] {
			fmt.Println("Нет")
			return
		}
	}
	fmt.Println("Палиндром")
}
