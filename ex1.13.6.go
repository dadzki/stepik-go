package main

import "fmt"

func main() {
	var a, b rune
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Printf("%v", float32(a+b)/2)
}
