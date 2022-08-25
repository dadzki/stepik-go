package main

import "fmt"

func main() {
	var n, val int
	var arr []int
	fmt.Scan(&n)

	if n < 4 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
	}

	fmt.Println(arr[3])
}
