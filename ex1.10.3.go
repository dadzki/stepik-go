package main

import "fmt"

func main() {
	var n, number int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if number >= 10 && number < 100 && number%8 == 0 {
			sum = sum + number
		}
	}
	fmt.Println(sum)
}
