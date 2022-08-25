package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(getFirst(a))
}

func getFirst(a int) int {
	if a < 10 {
		return a
	} else {
		return getFirst(a / 10)
	}
}
