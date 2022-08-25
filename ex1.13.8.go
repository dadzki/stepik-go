package main

import "fmt"

func main() {
	var n int
	var number, min int
	fmt.Scan(&n)

	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if i == 0 || number < min {
			count = 0
			min = number
		}
		if number == min {
			count++
		}
	}
	fmt.Println(count)
}
