package main

import (
	"fmt"
)

func main() {
	var a, fibPrev, fibNext, n int
	fmt.Scan(&a)

	if a <= 1 {
		return
	}

	fibPrev, fibNext = 0, 1
	n = 1
	for fibNext <= a {
		if fibNext == a {
			fmt.Println(n)
			return
		}
		fibPrev, fibNext = fibNext, fibPrev+fibNext
		n++
	}

	fmt.Println(-1)
}
