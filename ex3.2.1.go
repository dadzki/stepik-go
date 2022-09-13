package main

import (
	"fmt"
)

func convert(x int64) uint16 {
	return uint16(x)
}

func main() {
	var x int64

	fmt.Scan(&x)

	fmt.Println(convert(x))
}
