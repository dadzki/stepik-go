package main

import "fmt"

func main() {
	var n int
	var number int
	fmt.Scan(&n)

	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if number == 0 {
			count++
		}
	}
	fmt.Println(count)
}
