package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())

	fmt.Println(vote(0, 0, 1))
	fmt.Println(vote(1, 0, 1))

	fmt.Println(fibonacci(4))
}

func minimumFromFour() int {
	var n, min int
	for i := 0; i < 4; i++ {
		fmt.Scan(&n)
		if i == 0 || n < min {
			min = n
		}

	}
	return min
}

func vote(x int, y int, z int) int {
	if x+y+z > 1 {
		return 1
	} else {
		return 0
	}
}

func fibonacci(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		a, b = b, b+a
	}
	return a
}
