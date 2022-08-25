package main

import "fmt"

func main() {
	var n, val int
	var arr []int
	fmt.Scan(&n)

	if n < 1 && n > 100 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		arr = append(arr, val)
	}

	count := 0
	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			count++
		}
	}

	fmt.Println(count)
}
