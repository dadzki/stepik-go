package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func FindMaxInString(str string) (int, error) {
	rs := []rune(str)

	if len(rs) > 1000 {
		return 0, errors.New("ошибка ввода")
	}

	var maxValue int
	for i := 0; i < len(rs); i++ {
		if !unicode.IsDigit(rs[i]) {
			return 0, errors.New("ошибка ввода - строка должна содержать только цифры")
		}
		val, _ := strconv.Atoi(string(rs[i]))

		if i == 0 || maxValue < val {
			maxValue = val
		}
	}

	return maxValue, nil
}

func main() {
	var str string

	fmt.Scan(&str)

	max, err := FindMaxInString(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(max)
}
