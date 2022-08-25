package main

import (
	"fmt"
)

func main() {
	var x string
	fmt.Scan(&x)

	for i, char := range x {
		if i%2 != 0 {
			fmt.Print(string(char))
		}
	}
}
