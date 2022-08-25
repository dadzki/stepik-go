package main

import "fmt"

func main() {
	array := [5]int{}
	var a, max int

	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	for i := 0; i < len(array); i++ {
		if i == 0 {
			max = array[i]
		}
		if array[i] > max {
			max = array[i]
		}
	}

	fmt.Println(max)
}
