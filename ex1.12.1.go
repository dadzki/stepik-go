package main

import "fmt"

func main() {
	var workArray [10]byte
	var indexChangeArray [6]byte
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	for i := 0; i < 6; i++ {
		fmt.Scan(&indexChangeArray[i])
	}

	for i := 0; i < len(indexChangeArray); i++ {
		workArray[indexChangeArray[i]], workArray[indexChangeArray[i+1]] = workArray[indexChangeArray[i+1]], workArray[indexChangeArray[i]]
		i++
	}
	for i := 0; i < len(workArray); i++ {
		fmt.Printf("%v ", workArray[i])
	}
}
