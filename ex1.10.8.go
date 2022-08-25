package main

import "fmt"

func main() {
	var number1, number2, number2tmp, d1, d2 int

	fmt.Scan(&number1, &number2)
	number1 = revertNumber(number1)
	number2tmp = number2

	for number1 > 0 {
		d1 = number1 % 10
		number2 = number2tmp
		for number2 > 0 {
			d2 = number2 % 10
			if d1 == d2 {
				fmt.Print(d1)
				fmt.Print(" ")
				break
			}
			number2 = number2 / 10
		}
		number1 = number1 / 10
	}
}

func revertNumber(n int) int {
	m := 0
	for n > 0 {
		m = m*10 + n%10
		n = n / 10
	}
	return m
}
