package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	sum := 0
	if a < b && a <= 100 && b <= 100 {
		for i := a; i <= b; i++ {
			sum = sum + i
		}
		fmt.Println(sum)
	}
}
