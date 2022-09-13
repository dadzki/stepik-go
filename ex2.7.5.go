package main

import (
	"fmt"
	"math"
)

var k, p, v float64

func T() float64 {
	return 6 / W()
}

func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v
}

func main() {

	var s = []map[int]string{}
	var s []map[int]int
	m := make(map[float32]int)

	k, p, v = 1296, 6, 6
	fmt.Println(T())
}
