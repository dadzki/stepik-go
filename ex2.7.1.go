package main

import (
	"fmt"
	"math"
)

func GetGipotLength(a int, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	fmt.Println(GetGipotLength(a, b))
}
