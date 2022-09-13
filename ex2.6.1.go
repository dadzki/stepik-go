package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("ошибка")
		return 0, err
	}

	return a / b, nil
}

func main() {
	var a, b, res int

	fmt.Scan(&a, &b)

	res, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
