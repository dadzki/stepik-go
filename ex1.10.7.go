package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	year := 0
	for {
		year++
		x = x + x*p/100
		if x > y {
			break
		}
	}
	fmt.Println(year)
}
