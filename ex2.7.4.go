package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var str string

	fmt.Scan(&str)
	rs := []rune(str)

	for _, v := range rs {
		if !unicode.IsDigit(v) {
			panic("ошибка ввода - строка должна содержать только цифры")
		}
		val, _ := strconv.Atoi(string(v))
		fmt.Print(val * val)
	}

}
