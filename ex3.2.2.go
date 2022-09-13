package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func clearStringForDigits(str string) int64 {
	var newStr string
	runeStr := []rune(str)
	for _, v := range runeStr {
		if unicode.IsDigit(v) {
			newStr += string(v)
		}
	}

	res, err := strconv.ParseInt(newStr, 10, 64)
	if err != nil {
		panic("ошибка")
	}
	return res
}

func adding(x string, y string) int64 {
	return clearStringForDigits(x) + clearStringForDigits(y)
}

func main() {
	var x, y string

	fmt.Scan(&x, &y)

	fmt.Println(adding(x, y))
}
